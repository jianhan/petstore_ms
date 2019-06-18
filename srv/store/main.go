// Package main initialize service and handlers, prepare dependencies other packages need
// such as db connection, handlers, etc.. It is entry point of store service.
package main

import (
	_ "github.com/go-sql-driver/mysql"
	handler "github.com/jianhan/petstore_ms/srv/store/handler"
	"github.com/jianhan/petstore_ms/srv/store/mysql"
	pet "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// initialize service
	service := micro.NewService(
		micro.Name("go.micro.srv.store"),
		micro.Version("v1"),
	)
	service.Init()

	// get db connection
	db, err := mysql.Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// register handlers
	pet.RegisterPetServiceHandler(service.Server(), handler.NewPetServiceHandler(mysql.NewPetDataStore(db)))

	// run service
	if err := service.Run(); err != nil {
		panic(err)
	}
}
