package app

import (
	"io"
	"net/http"
	"strings"
)

func (app *Application) reverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	var backendURL, backendPath string
	switch {
	case strings.HasPrefix(path, "/s1/"):
		backendURL = "http://localhost:4200"
		backendPath = strings.TrimPrefix(path, "/s1")
	case strings.HasPrefix(path, "/s2/"):
		backendURL = "http://localhost:2200"
		backendPath = strings.TrimPrefix(path, "/s2")
	default:
		http.Error(w, "URL Not Found", http.StatusNotFound)
		return
	}

	fullURL := backendURL + backendPath

	app.Logger.Info("Forwarding request", "method", r.Method, "path", path, "backend", backendURL)

	if r.Method == http.MethodGet {
		if cachedResp, found := app.Cache.Get(path); found {
			w.WriteHeader(http.StatusOK)
			w.Write(cachedResp)
			app.Logger.Info("Cache hit", "path", path)
			return
		}
	}

	req, err := http.NewRequest(r.Method, fullURL, r.Body)
	if err != nil {
		app.Logger.Error("Failed to create request", "error", err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		app.Logger.Error("Error forwarding request", "error", err)
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
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

	if r.Method == http.MethodGet && resp.StatusCode == http.StatusOK {
		app.Cache.Store(path, bodyBytes)
		app.Logger.Info("Response cached", "path", path)
	}
}
