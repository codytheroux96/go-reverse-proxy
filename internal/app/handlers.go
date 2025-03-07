package app

import (
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body on server one")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *Application) HandleServerOnePost (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:4200/s1list")
	if err != nil {
		fmt.Println("not hitting server one")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body on server one")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *Application) HandleServerTwoGet (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:2200/s2health")
	if err != nil {
		fmt.Println("not hitting server two")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body on server two")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *Application) HandleServerTwoPost (w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:2200/s2list")
	if err != nil {
		fmt.Println("not hitting server two")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body on server two")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}