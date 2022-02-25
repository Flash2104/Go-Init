package main

import "fmt"

type Supplier struct {
	name, city string
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

type Person struct {
	name, city string
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		p.printDetails()
	}
	(*Product).printDetails(&Product{"Kayak", "Watersports", 275})
	fmt.Println("----------------------------")
	fmt.Println("Understanding Method Overloading")

	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50, []string{}}
	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:",
		insurance.monthlyFee*float64(insurance.durationMonths))

	// expenses := []Expense{
	// 	Product{"Kayak", "Watersports", 275},
	// 	Service{"Boat Cover", 12, 89.50},
	// }
	// for _, expense := range expenses {
	// 	fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(expenses))

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50, []string{}},
		},
	}
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))

	fmt.Println("--------------------------")
	fmt.Println("Understanding the Effect of Pointer Method Receivers")
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
	product.price = 100
	fmt.Println("Expense method changed result:", expense.getCost(false))

	fmt.Println("--------------------------")
	fmt.Println("Comparing Interface Values")
	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	// var e3 Expense = Service{description: "Boat Cover"}
	// var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	// fmt.Println("e3 == e4", e3 == e4)

	fmt.Println("--------------------------")
	fmt.Println("Type Assertion")
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		Product{"Product", "12", 8},
	}
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}

	fmt.Println("--------------------------")
	fmt.Println("Switching on Dynamic Types")

	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
			break
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}

	fmt.Println("--------------------------")
	fmt.Println("Using the Empty Interface")
	expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
	fmt.Println("--------------------------")
	fmt.Println("Using the Empty Interface for Function Parameters")
	for _, item := range data {
		processItem(item)
	}

	processItem(data...)
}

func processItem(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}

}
