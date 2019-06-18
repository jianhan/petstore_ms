package store

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func (r InsertPetRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Pet, validation.Required, validation.NotNil),
	)

	return nil
}

func (p Pet) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Status, validation.In(Available.String(), Pending.String(), Sold.String())),
	)

	return nil
}
