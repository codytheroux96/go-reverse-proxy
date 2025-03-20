package app

import (
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

}
