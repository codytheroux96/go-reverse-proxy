package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (app *Application) determineBackendURl(requestPath string) (string, error) {
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

func (app *Application) HandleGetRequest(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) HandlePostRequest(w http.ResponseWriter, r *http.Request) {

}
