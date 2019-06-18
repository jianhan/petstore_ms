// Package main initialize service and handlers, prepare dependencies other packages need
// such as db connection, handlers, etc.. It is entry point of store service.
package main

import (
	handler "github.com/jianhan/petstore_ms/srv/store/handler"
	pet "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	"github.com/micro/go-micro"
)

func main() {
	// initialize service
	service := micro.NewService(
		micro.Name("go.micro.srv.store"),
		micro.Version("v1"),
	)
	service.Init()

	// register handlers
	pet.RegisterPetServiceHandler(service.Server(), handler.NewPetServiceHandler())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
