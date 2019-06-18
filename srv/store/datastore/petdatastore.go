package datastore

import store "github.com/jianhan/petstore_ms/srv/store/proto/pet"

type PetDataStore interface {
	UpsertPets([]*store.Pet) error
}
