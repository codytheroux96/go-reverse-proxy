package app

import (
	"net/http"
)

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/s1get", app.HandleServerOneGet)
	mux.HandleFunc("/s2get", app.HandleServerTwoGet)

	mux.HandleFunc("/s1post", app.HandleServerOnePost)
	mux.HandleFunc("/s2post", app.HandleServerTwoPost)

	return mux
}
