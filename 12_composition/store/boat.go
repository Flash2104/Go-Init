package store

// Boat Product struct
type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

// NewBoat constructor
func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{NewProduct(name, "Watersports", price), capacity, motorized}
}
