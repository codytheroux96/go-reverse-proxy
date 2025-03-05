package main

import (
	"github.com/codytheroux96/go-reverse-proxy/test_servers"
	"github.com/codytheroux96/go-reverse-proxy/internal/app"
)

func main() {
	app := app.NewApplication()
	app.Logger.Info("MESSAGE FROM MAIN SERVER: APPLICATION IS RUNNING!!!")

	go func() {
		testservers.ServerOne()
	}()

	go func() {
		testservers.ServerTwo()
	}()

	select{}
}
