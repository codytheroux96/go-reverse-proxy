package server_one

import (
	"net/http"
	"os"
	"time"
)

func ServerOne() {
	app := newApplication()

	serverOne := &http.Server{
		Addr:         ":4200",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Info("SERVER ONE IS RUNNING", "addr", serverOne.Addr)

	err := serverOne.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
