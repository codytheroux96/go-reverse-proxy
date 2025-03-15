package app

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type RateLimiterConfig struct {
	enabled bool
	rps     float64
	burst   int
}

type Application struct {
	Logger *slog.Logger
	Cache  *ResponseCache
	config struct {
		Limiter RateLimiterConfig
	}
}

func NewApplication() *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		Logger: logger,
		Cache:  NewResponseCache(30 * time.Second),
	}

	go app.Cache.Cleanup(15 * time.Second)

	app.config.Limiter = RateLimiterConfig{
		enabled: true,
		rps:     50,
		burst:   250,
	}

	return app
}

func (app *Application) LogRequest(r *http.Request) {
	app.Logger.Info("Incoming Request", "method", r.Method, "path", r.URL.Path)
}