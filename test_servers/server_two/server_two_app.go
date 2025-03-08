package server_two

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type application struct {
	logger *slog.Logger
}

func newApplication() *application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &application{
		logger: logger,
	}
}

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/s2health", app.handleHealthcheck)
	r.Post("/s2list", app.handleList)

	return r
}

func (app *application) handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Two Is Up And Healthy\n"))
}

func (app *application) handleList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Two Is Serving A Fake List\n"))
}
