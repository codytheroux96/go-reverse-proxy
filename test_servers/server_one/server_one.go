package server_one

import (
	"flag"
	"net/http"
	"os"
	"time"
)

func ServerOne() {
	addr := flag.String("addr", ":4200", "HTTP Network Address")
	flag.Parse()

	app := newApplication()

	serverOne := &http.Server{
		Addr:         *addr,
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
