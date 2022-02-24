package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

type StockLevel struct {
	Product
	Alternate Product
	count     int
}

func main() {
	acme := &Supplier{"Acme Co", "New York"}
	stockItem := StockLevel{
		Product{name: "Kayak", category: "Watersports", price: 275.00},
		Product{"Lifejacket", "Watersports", 48.95, acme},
		100,
	}
	kayak := stockItem.Product
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Price:", stockItem.Product.price)
	fmt.Println("Price:", stockItem.Alternate.price)
	fmt.Println("Price:", stockItem.price)
	fmt.Println("Count:", stockItem.count)

	fmt.Println("-------------------------")
	fmt.Println("Comparint struct values")
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
	type Item struct {
		name     string
		category string
		price    float64
		*Supplier
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00, Supplier: acme}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00, Supplier: acme}
	fmt.Println("prod == item:", prod == Product(item))

	fmt.Println("-----------------------------")
	fmt.Println("Defining Anonymous Struct Types")
	writeName(prod)
	writeName(item)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println("Builder String:", builder.String())

	fmt.Println("-------------------------")
	fmt.Println("Creating Arrays, Slices, and Maps Containing Struct Values")
	array := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)
	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00, acme},
			Alternate: Product{"Lifejacket", "Watersports", 48.95, acme},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)

	fmt.Println("-------------------------")
	fmt.Println("Understanding Structs and Pointers")
	p5 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p6 := p5
	p7 := &p5
	p5.name = "Original Kayak"
	fmt.Println("P5:", p5.name)
	fmt.Println("P6:", p6.name)
	fmt.Println("P6:", p7.name)
	kayak1 := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}

	calcTax(kayak1)
	fmt.Println("Name:", kayak1.name, "Category:",
		kayak1.category, "Price", kayak1.price)

	fmt.Println("-------------------------")
	fmt.Println("Understanding Struct Constructor Functions")
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}

	p8 := newProduct("Kayak", "Watersports", 275, acme)
	p9 := *p8
	p8.name = "Original Kayak"
	p8.Supplier.name = "BoatCo"
	for _, p := range []Product{*p8, p9} {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}

	fmt.Println("-------------------------")
	fmt.Println("Understanding Zero Value for Structs and Pointers to Structs")
	var prod1 Product
	var prodPtr *Product
	fmt.Println("Value:", prod1.name, prod1.category, prod1.price, prod1.Supplier.name)
	fmt.Println("Pointer:", prodPtr)
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price, supplier}
}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func writeName(val struct {
	name, category string
	price          float64
	*Supplier
}) {
	fmt.Println("Name:", val.name)
}
