package app

import (
	"bytes"
	"io"
	"net/http"
)

func (app *Application) HandleServerOneGet(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	resp, err := http.Get("http://localhost:4200/s1health")
	if err != nil {
		app.Logger.Error("error when making GET request to server one", "error", err)
		http.Error(w, "error when making GET request to server one", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		app.Logger.Error("could not copy body from server one on GET request", "error", err)
		http.Error(w, "could not copy body from server one on GET request", http.StatusInternalServerError)
		return
	}
}

func (app *Application) HandleServerOnePost(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4200/s1list", bytes.NewBuffer([]byte{}))
	if err != nil {
		app.Logger.Error("error when making a POST request to server one", "error", err)
		http.Error(w, "error when making a POST request to server one", http.StatusInternalServerError)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		app.Logger.Error("error when receiving response from server one on POST request", "error", err)
		http.Error(w, "error when receiving response from server one on POST request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		app.Logger.Info("could not copy body from server one on POST request", "error", err)
		http.Error(w, "could not copy body from server one on POST request", http.StatusInternalServerError)
	}
}

func (app *Application) HandleServerTwoGet(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	resp, err := http.Get("http://localhost:2200/s2health")
	if err != nil {
		app.Logger.Info("error when making GET request to server two", "error", err)
		http.Error(w, "error when making GET request to server two", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		app.Logger.Error("could not copy body from server two on GET request", "error", err)
		http.Error(w, "could not copy body from server two on GET request", http.StatusInternalServerError)
	}
}

func (app *Application) HandleServerTwoPost(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:2200/s2list", bytes.NewBuffer([]byte{}))
	if err != nil {
		app.Logger.Info("error when making POST request to server two", "error", err)
		http.Error(w, "error when making POST request to server two", http.StatusInternalServerError)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		app.Logger.Error("error when receiving response from server two on POST request", "error", err)
		http.Error(w, "error when receiving response from server two on POST request", http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		app.Logger.Error("could not copy body from server two on POST request", "error", err)
		http.Error(w, "could not copy body from server two on POST request", http.StatusInternalServerError)
	}
}
