package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/jianhan/petstore_ms/srv/api/handler"
	"github.com/jianhan/petstore_ms/srv/api/route"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// main is entry point to launch the api server.
func main() {

	// load env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file")
	}

	// initialize rpc service in order to inject into pet handler
	client := micro.NewService(micro.Name("go.micro.srv.store.client"))
	client.Init()
	petService := store.NewPetService("go.micro.srv.store", client.Client())

	// initialize pet handler
	petHandler := handler.NewPetHandler(petService)

	// initialize routes:w
	router := route.InitRoutes(petHandler)

	// get classic middleware from negroni
	n := negroni.Classic()
	n.UseHandler(router)
	srv := &http.Server{
		Handler:      handlers.CORS()(n),
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	// launch server
	logrus.Fatal(srv.ListenAndServe())
}
