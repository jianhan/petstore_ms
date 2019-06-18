package handler

import (
	"context"

	datastore "github.com/jianhan/petstore_ms/srv/store/datastore"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

type petHandler struct {
	petDataStore datastore.PetDataStore
}

func (p *petHandler) InsertPet(ctx context.Context, req *store.InsertPetRequest, rsp *store.InsertPetResponse) error {
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

func (p *petHandler) UpdatePet(ctx context.Context, req *store.UpdatePetRequest, rsp *store.UpdatePetResponse) (err error) {
	if rsp.RowsAffected, err = p.petDataStore.UpdatePet(req); err != nil {
		return err
	}

	return nil
}

func NewPetServiceHandler(petDataStore datastore.PetDataStore) store.PetServiceHandler {
	return &petHandler{petDataStore: petDataStore}
}
