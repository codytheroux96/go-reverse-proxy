package testservers

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/codytheroux96/go-reverse-proxy/internal/app"
)

func ServerOne() {
	addr := flag.String("addr", ":4200", "HTTP Network Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &app.Application{
		Logger: logger,
	}

	serverOne := &http.Server{
		Addr: *addr,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	app.Logger.Info("starting server one", "addr", serverOne.Addr)

	err := serverOne.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}