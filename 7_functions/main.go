package main

import "fmt"

func main() {
	fmt.Println("====================")
	fmt.Println("Variadic Parameters")
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Kayak", names...)
	printSuppliers("Lifejacket", "Sail Safe Co")
	printSuppliers("Soccer Ball")

	fmt.Println("====================")
	fmt.Println("Pointers as func params")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function", val1, val2)

	fmt.Println("====================")
	fmt.Println("Using Function Results")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		priceWithTax, _ := calcTax(price)
		fmt.Println("Product: ", product, "Price:", priceWithTax)
	}

	fmt.Println("====================")
	fmt.Println("Using Named Results and defer")
	total1, tax1 := calcTotalPrice(products)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
}

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}
