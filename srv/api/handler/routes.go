package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(petHandler petHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v2").Subrouter()

	// define routes
	api.HandleFunc("/pet", petHandler.insertPet).Methods(http.MethodPost)
	api.HandleFunc("/pet/{petId}", petHandler.updatePet).Methods(http.MethodPost)

	return router
}
