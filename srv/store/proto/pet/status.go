package store

// PetStatus defines constaints for pet status.
type PetStatus int

const (
	// Available indicates pet are available for sale.
	Available PetStatus = iota

	// Pending indicates pet is pending.
	Pending

	// Sold indicates pet already sold.
	Sold
)

// String returns string representation of pet status.
func (p PetStatus) String() string {
	return [...]string{"available", "pending", "sold"}[p]
}
