package main

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}

	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
		b.Name = "New" + b.Name
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, false, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	type OfferBundle struct {
		*store.Product
	}

	bundle := OfferBundle{
		store.NewProduct("Lifrejacket", "Watersports", 48.95),
	}
	fmt.Println("Price:", bundle.Price(0))

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for k, v := range products {
		switch item := v.(type) {
		case store.Describable:
			fmt.Println("Describable Name:", item.GetName(), "Category:", item.GetCategory(), "Price:", item.(store.ItemForSale).Price(0.2))
		case *store.Product:
			fmt.Println("Switch-case. Product Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Switch-case. Boat Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))

		default:
			fmt.Println("Key:", k, "Price:", v.Price(0.2))
		}

	}
}
