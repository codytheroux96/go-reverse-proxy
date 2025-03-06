package app

import (
	"log/slog"
	"os"
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
