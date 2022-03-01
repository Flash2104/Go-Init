package main

import (
	"fmt"
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
	// "github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	// color.Green("Name: " + cart.CustomerName)
	// color.Cyan("Total: " + currencyFmt.ToCurrency(cart.GetTotal()))
	fmt.Println("Name: ", cart.CustomerName)
	fmt.Println("Total: ", currencyFmt.ToCurrency(cart.GetTotal()))
}
