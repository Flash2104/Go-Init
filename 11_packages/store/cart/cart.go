package cart

import "packages/store"

// Cart struct info
type Cart struct {
	CustomerName string
	Products     []store.Product
}

// GetTotal method for get sum of all products prices
func (cart *Cart) GetTotal() (total float64) {
	for _, p := range cart.Products {
		total += p.Price()
	}
	return
}
