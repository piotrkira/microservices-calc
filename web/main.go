package main

import (
	server "github.com/piotrkira/microservices-calc/web/server"
)

func main() {
	s := server.New()
	s.Start()
}
