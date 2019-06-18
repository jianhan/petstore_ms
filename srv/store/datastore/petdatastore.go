package datastore

import store "github.com/jianhan/petstore_ms/srv/store/proto/pet"

type PetDataStore interface {
	InsertPet(*store.Pet) (int64, error)
	UpdatePet(*store.UpdatePetRequest) (int64, error)
	FindPetById(int64) (*store.Pet, error)
}
