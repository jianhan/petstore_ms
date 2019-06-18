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
	// p.petDataStore.UpsertPets(req.Pets)
	rsp.Pets = nil
	return nil
}

func NewPetServiceHandler() store.PetServiceHandler {
	return &petHandler{}
}
