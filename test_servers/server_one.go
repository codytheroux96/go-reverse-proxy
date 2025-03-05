package testservers

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func ServerOne() {
	addr := flag.String("addr", ":4200", "HTTP Network Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	serverOne := &http.Server{
		Addr: *addr,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	logger.Info("starting server one", "addr", serverOne.Addr)

	err := serverOne.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}