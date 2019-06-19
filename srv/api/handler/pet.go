package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jianhan/petstore_ms/srv/api/response"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

// PetHandler defines method that pet handler will implement.
type PetHandler interface {
	InsertPet(w http.ResponseWriter, r *http.Request)
	UpdatePet(w http.ResponseWriter, r *http.Request)
}

// petHandler implements PetHandler interface to process all pet related HTTP requests.
type petHandler struct {
	petService store.PetService
}

// InsertPet inserts a new pet via pet micro service.
func (p *petHandler) InsertPet(w http.ResponseWriter, r *http.Request) {

	// decode request
	decoder := json.NewDecoder(r.Body)
	var insertRequest *store.InsertPetRequest
	if err := decoder.Decode(&insertRequest); err != nil {
		response.WithJSONStr(w, err.Error(), http.StatusBadRequest)
		return
	}

	// process insert pet via pet micro service of InsertPet method.
	rsp, err := p.petService.InsertPet(r.Context(), insertRequest)
	if err != nil {
		response.WithJSONStr(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WithJSONData(w, rsp.Pet, http.StatusOK)
}

// UpdatePet handle update pet http call.
func (p *petHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {

	// decode request
	decoder := json.NewDecoder(r.Body)
	var updateRequest *store.UpdatePetRequest
	if err := decoder.Decode(&updateRequest); err != nil {
		response.WithJSONStr(w, err.Error(), http.StatusBadRequest)
		return
	}

	// process update pet via pet micro service with UpdatePet method.
	if _, err := p.petService.UpdatePet(r.Context(), updateRequest); err != nil {
		response.WithJSONStr(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WithJSONBytes(w, []byte{}, http.StatusOK)
}

// NewPetHandler returns petHandler which implements PetHandler.
func NewPetHandler(petService store.PetService) PetHandler {
	return &petHandler{petService: petService}
}
