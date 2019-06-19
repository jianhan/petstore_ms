package store

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate validates InsertPetRequest.
func (r InsertPetRequest) Validate() error {

	// within request pet must not be nil
	return validation.ValidateStruct(&r,
		validation.Field(&r.Pet, validation.Required, validation.NotNil),
	)
}

// Validate validates pet.
func (p Pet) Validate() error {

	// pet must contain name and status.
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Status, validation.Required, validation.In(Available.String(), Pending.String(), Sold.String())),
	)
}

// Validate validates  UpdatePetRequest.
func (r UpdatePetRequest) Validate() error {

	// when updating pet, id , status and name must be required. also status must be a valid one.
	return validation.ValidateStruct(&r,
		validation.Field(&r.Id, validation.Required),
		validation.Field(&r.Status, validation.Required, validation.In(Available.String(), Pending.String(), Sold.String())),
		validation.Field(&r.Name, validation.Required),
	)
}
