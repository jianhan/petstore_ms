package handler

import (
	"context"

	datastore "github.com/jianhan/petstore_ms/srv/store/datastore"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

type petHandler struct {
	petDataStore datastore.PetDataStore
}

func (p *petHandler) UpsertPets(ctx context.Context, req *store.UpsertPetsRequest, rsp *store.UpsertPetsResponse) error {
	if err := p.petDataStore.UpsertPets(req.Pets); err != nil {
		return err
	}

	return nil
}

func NewPetServiceHandler(petDataStore datastore.PetDataStore) store.PetServiceHandler {
	return &petHandler{petDataStore: petDataStore}
}
