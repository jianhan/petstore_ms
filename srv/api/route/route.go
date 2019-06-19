// Package route register all handler for routes.
package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jianhan/petstore_ms/srv/api/handler"
)

// InitRoutes initialize routes for handlers.
func InitRoutes(petHandler handler.PetHandler) *mux.Router {

	// define router
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v2").Subrouter()

	// define routes
	api.HandleFunc("/pet", petHandler.InsertPet).Methods(http.MethodPost)
	api.HandleFunc("/pet/{petId}", petHandler.UpdatePet).Methods(http.MethodPost)

	return router
}
