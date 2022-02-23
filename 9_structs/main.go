package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		count int
	}
	stockItem := StockLevel{
		Product{"Kayak", "Watersports", 275.00},
		100,
	}
	kayak := stockItem.Product
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Price:", stockItem.Product.price)
	fmt.Println("Count:", stockItem.count)
}
