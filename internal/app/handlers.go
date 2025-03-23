package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (app *Application) reverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.HandleGetRequest(w, r)
	case http.MethodPost:
		app.HandlePostRequest(w, r)
	default:
		http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
	}
}

func (app *Application) determineBackendURL(requestPath string) (string, error) {
	var backendURL string

	switch {
	case strings.HasPrefix(requestPath, "/s1/"):
		backendURL = "http://localhost:4200" + strings.TrimPrefix(requestPath, "/s1")
	case strings.HasPrefix(requestPath, "/s2/"):
		backendURL = "http://localhost:2200" + strings.TrimPrefix(requestPath, "/s2")
	default:
		app.Logger.Error("Invalid backend path", "path", requestPath)
		return "", fmt.Errorf("no matching backend for path: %s", requestPath)
	}

	return backendURL, nil
}

func (app *Application) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if cachedResp, found := app.Cache.Get(path); found {
		w.WriteHeader(http.StatusOK)
		w.Write(cachedResp)
		app.Logger.Info("Cache hit", "path", path)
		return
	}

	backendURL, err := app.determineBackendURL(path)
	if err != nil {
		http.Error(w, "invalid backend path", http.StatusNotFound)
		return
	}

	maxRetries := 3
	backoffTimes := []time.Duration{100 * time.Millisecond, 500 * time.Millisecond, 2 * time.Second}

	var resp *http.Response
	for attempt := 1; attempt <= maxRetries; attempt++ {
		app.Logger.Info("Forwarding GET request", "url", backendURL, "attempt", attempt)
		resp, err = http.Get(backendURL)

		if err != nil || (resp.StatusCode >= 500 && resp.StatusCode <= 504) {
			app.Logger.Info("Retrying GET request", "url", backendURL, "attempt", attempt)
			if attempt < maxRetries {
				time.Sleep(backoffTimes[attempt-1])
			}
			continue
		}
		break
	}

	if err != nil || resp.StatusCode >= 500 {
		app.Logger.Error("Final failure on GET request", "url", backendURL, "response_code", resp.StatusCode)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		app.Logger.Error("Failed to read response body", "error", err)
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		app.Cache.Store(path, bodyBytes)
		app.Logger.Info("Response cached", "path", path)
	}

}

func (app *Application) HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	backendURL, err := app.determineBackendURL(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid backend path", http.StatusNotFound)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		app.Logger.Error("failed to read request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	maxRetries := 3
	backoffTimes := []time.Duration{100 * time.Millisecond, 500 * time.Millisecond, 2 * time.Second}

	var resp *http.Response
	for attempt := 1; attempt <= maxRetries; attempt++ {
		bodyReader := bytes.NewReader(bodyBytes)

		req, err := http.NewRequest(http.MethodPost, backendURL, bodyReader)
		if err != nil {
			app.Logger.Error("failed to create POST request", "error", err)
			http.Error(w, "failed to create request", http.StatusInternalServerError)
			return
		}

		for key, values := range r.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil || (resp.StatusCode >= 500 && resp.StatusCode <= 504) {
			app.Logger.Info("retrying POST request", "url", backendURL, "attempt", attempt)
			if attempt < maxRetries {
				time.Sleep(backoffTimes[attempt-1])
			}
			continue
		}
		break
	}

	if resp == nil {
		app.Logger.Error("POST request failed - no response received", "url", backendURL, "error", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	if resp.StatusCode >= 500 {
		app.Logger.Error("POST request failed - backend error", "url", backendURL, "status", resp.StatusCode)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
