package app

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Application struct {
	Logger *slog.Logger
	// collection of servers to handle/serve will go here (server1, server2, etc)
}

func NewApplication() *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return &Application{
		Logger: logger,
	}
}

func (app *Application) LogRequest(r *http.Request) {
	app.Logger.Info("Incoming Request", "method", r.Method, "path", r.URL.Path, "timestamp", time.Now().Format(time.RFC3339))
}
