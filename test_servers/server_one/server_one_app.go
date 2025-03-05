package server_one

import (
	"log/slog"
	"net/http"
	"os"
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
	mux := http.NewServeMux()

	mux.HandleFunc("/health", app.healthCheckHandler)
	mux.HandleFunc("/list", app.listHandler)

	return mux
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server One Is Up And Healthy"))
}

func (app *application) listHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server One Is Serving A Fake List"))
}
