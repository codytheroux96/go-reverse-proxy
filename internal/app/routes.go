package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/s1get", app.HandleServerOneGet)
	r.Get("/s2get", app.HandleServerTwoGet)

	r.Post("/s1post", app.HandleServerOnePost)
	r.Post("/s2post", app.HandleServerTwoPost)

	return r
}
