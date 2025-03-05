package main

import (
	"github.com/codytheroux96/go-reverse-proxy/internal/app"
	"github.com/codytheroux96/go-reverse-proxy/test_servers/server_one"
	"github.com/codytheroux96/go-reverse-proxy/test_servers/server_two"
)

func main() {
	app := app.NewApplication()
	app.Logger.Info("MESSAGE FROM MAIN SERVER: APPLICATION IS RUNNING!!!")

	go func() {
		server_one.ServerOne()
	}()

	go func() {
		server_two.ServerTwo()
	}()

	select{}
}
