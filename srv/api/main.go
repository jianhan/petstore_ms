package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jianhan/petstore_ms/srv/api/handler"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/urfave/negroni"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	client := micro.NewService(micro.Name("go.micro.srv.store.client"))
	client.Init()

	petService := store.NewPetService("go.micro.srv.store", client.Client())
	petHandler := handler.NewPetHandler(petService)

	router := handler.InitRoutes(petHandler)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	// Or use a middleware with the Use() function

	n.UseHandler(router)
	srv := &http.Server{
		Handler:      n,
		Addr:         fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
