package store

type PetStatus int

const (
	Available PetStatus = iota
	Pending
	Sold
)

func (p PetStatus) String() string {
	return AllPetStatus()[p]
}

func AllPetStatus() [3]string {
	return [...]string{"available", "pending", "sold"}
}
