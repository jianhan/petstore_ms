package handler

import (
	"encoding/json"
	"net/http"

	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

type petHandler struct {
	petService store.PetService
}

func (p *petHandler) insertPet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var insertRequest *store.InsertPetRequest
	if err := decoder.Decode(&insertRequest); err != nil {
		errorWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	rsp, err := p.petService.InsertPet(r.Context(), insertRequest)
	if err != nil {
		errorWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	petBytes, err := json.Marshal(&rsp.Pet)
	if err != nil {
		errorWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWithJSON(w, petBytes, http.StatusOK)
}

func (p *petHandler) updatePet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var updateRequest *store.UpdatePetRequest
	if err := decoder.Decode(&updateRequest); err != nil {
		errorWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := p.petService.UpdatePet(r.Context(), updateRequest); err != nil {
		errorWithJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseWithJSON(w, []byte{}, http.StatusOK)
}

func NewPetHandler(petService store.PetService) petHandler {
	return petHandler{petService: petService}
}
