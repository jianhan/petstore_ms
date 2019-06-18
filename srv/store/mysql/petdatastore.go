package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/jianhan/petstore_ms/srv/store/datastore"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

func NewPetDataStore(db *sql.DB) datastore.PetDataStore {
	return &petDataStore{db: db}
}

type petDataStore struct {
	db *sql.DB
}

func (p *petDataStore) InsertPet(pet *store.Pet) (int64, error) {
	photoUrlsStr, err := json.Marshal(pet.PhotoUrls)
	if err != nil {
		return 0, err
	}

	r, err := p.db.Exec("INSERT INTO pets (name, photo_urls, status) VALUES(?, ?, ?)", pet.Name, photoUrlsStr, pet.Status)
	if err != nil {
		return 0, err
	}

	// get last insert id
	lastInsertId, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (p *petDataStore) FindPetById(id int64) (*store.Pet, error) {
	pet := store.Pet{}
	photoUrlsStr := ""
	err := p.db.QueryRow("SELECT id, name, photo_urls, status from pets WHERE id = ?", id).Scan(&pet.Id, &pet.Name, &photoUrlsStr, &pet.Status)
	if err != nil {
		return nil, err
	}

	// convert json string to slice for photo urls
	if err := json.Unmarshal([]byte(photoUrlsStr), &pet.PhotoUrls); err != nil {
		return nil, err
	}

	return &pet, nil
}

func (p *petDataStore) UpdatePet(pet *store.UpdatePetRequest) (int64, error) {
	r, err := p.db.Exec("UPDATE pets SET name=?, status=? WHERE id = ?", pet.Name, pet.Status, pet.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
