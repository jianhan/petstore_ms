package mysql

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jianhan/petstore_ms/srv/store/datastore"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

func NewPetDataStore(db *sql.DB) datastore.PetDataStore {
	return &petDataStore{db: db}
}

type petDataStore struct {
	db *sql.DB
}

func (p *petDataStore) UpsertPets(pets []*store.Pet) error {
	for _, pet := range pets {
		if pet.Id > 0 {
			// update
			if err := p.updatePet(pet); err != nil {
				return err
			}
		} else {
			// insert
			if err := p.insertPet(pet); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *petDataStore) insertPet(pet *store.Pet) error {
	if _, err := p.db.Exec("INSERT INTO pets (name, photo_urls, status) VALUES(?, ?, ?)", pet.Name, pet.PhotoUrls, pet.Status); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *petDataStore) updatePet(pet *store.Pet) error {
	if _, err := p.db.Exec("UPDATE pets SET name=?, photo_urls=?, status=? WHERE id = ?", pet.Name, pet.PhotoUrls, pet.Status, pet.Id); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
