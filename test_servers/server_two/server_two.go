package server_two

import (
	"net/http"
	"os"
	"time"
)

func Serve() {
	app := newApplication()

	serverTwo := &http.Server{
		Addr:         ":2200",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Info("SERVER TWO IS RUNNING", "addr", serverTwo.Addr)

	if err := serverTwo.ListenAndServe(); err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
