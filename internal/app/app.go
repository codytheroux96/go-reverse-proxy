package app

import (
	"log/slog"
	"os"
)

type Application struct {
	Logger *slog.Logger
}

func NewApplication() *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	return &Application{
		Logger: logger,
	}
}