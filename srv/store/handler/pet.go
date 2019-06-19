package handler

import (
	"context"

	datastore "github.com/jianhan/petstore_ms/srv/store/datastore"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

// petHandler implements datastore.PetDataStore.
type petHandler struct {
	petDataStore datastore.PetDataStore
}

// InsertPet is RPC end point for handling inserting pet request.
func (p *petHandler) InsertPet(ctx context.Context, req *store.InsertPetRequest, rsp *store.InsertPetResponse) error {

	// validate request
	if err := req.Validate(); err != nil {
		return err
	}

	// validate pet
	if err := req.Pet.Validate(); err != nil {
		return err
	}

	// insert via datastore
	lastInsertID, err := p.petDataStore.InsertPet(req.Pet)
	if err != nil {
		return err
	}

	// fetch pet
	pet, err := p.petDataStore.FindPetById(lastInsertID)
	if err != nil {
		return err
	}

	// construct response
	rsp.Pet = pet

	return nil
}

// UpdatePet is RPC end point for handling UpdatePet request.
func (p *petHandler) UpdatePet(ctx context.Context, req *store.UpdatePetRequest, rsp *store.UpdatePetResponse) (err error) {

	if err := req.Validate(); err != nil {
		return err
	}

	if rsp.RowsAffected, err = p.petDataStore.UpdatePet(req); err != nil {
		return err
	}

	return nil
}

// NewPetServiceHandler returns a new pet handler.
func NewPetServiceHandler(petDataStore datastore.PetDataStore) store.PetServiceHandler {
	return &petHandler{petDataStore: petDataStore}
}
