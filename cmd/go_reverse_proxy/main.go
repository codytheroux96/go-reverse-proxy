package main

import (
	"github.com/codytheroux96/go-reverse-proxy/test_servers"
)

func main() {
	go func() {
		testservers.ServerOne()
	}()

	go func() {
		testservers.ServerTwo()
	}()

	select{}
}
