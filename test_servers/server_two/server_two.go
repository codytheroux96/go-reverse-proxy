package server_two

import (
	"net/http"
	"os"
	"time"
)

func ServerTwo() {
	app := newApplication()

	serverTwo := &http.Server{
		Addr:         ":2200",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Info("SERVER TWO IS RUNNING", "addr", serverTwo.Addr)

	err := serverTwo.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
