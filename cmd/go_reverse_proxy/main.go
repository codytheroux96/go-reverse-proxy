package main

import (
	"net/http"
	"os"
	"time"

	"github.com/codytheroux96/go-reverse-proxy/internal/app"
	"github.com/codytheroux96/go-reverse-proxy/test_servers/server_one"
	"github.com/codytheroux96/go-reverse-proxy/test_servers/server_two"
)

func main() {
	app := app.NewApplication()
	app.Logger.Info("MESSAGE FROM MAIN SERVER: APPLICATION IS RUNNING!!!")

	proxyServer := &http.Server{
		Addr:         ":8080",
		Handler:      app.RateLimit(app.Routes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		server_one.ServerOne()
	}()

	go func() {
		server_two.ServerTwo()
	}()

	err := proxyServer.ListenAndServe()
	if err != nil {
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
	select {}
}
