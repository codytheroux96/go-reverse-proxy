package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (app *Application) HandleServerOneGet (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:4200/s1health")
	if err != nil {
		fmt.Println("not hitting server one")
		os.Exit(1)
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
		fmt.Println("could not read body on server one")
	}
}

func (app *Application) HandleServerOnePost (w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:4200/s1list", bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("not hitting server one")
		os.Exit(1)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("not hitting server two")
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
		fmt.Println("could not read body on server one")
	}
}

func (app *Application) HandleServerTwoGet (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:2200/s2health")
	if err != nil {
		fmt.Println("not hitting server two")
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
		fmt.Println("could not read body on server two")
	}
}

func (app *Application) HandleServerTwoPost (w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:2200/s2list", bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("not hitting server two")
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("not hitting server two")
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
		fmt.Println("could not read body on server two")
	}
}