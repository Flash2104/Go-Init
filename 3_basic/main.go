package main

import (
	"fmt"
)

func main1() {
	price, tax, _ := 20.5, 2.05, true
	const quantity = 2
	fmt.Println("Total:", quantity*(price+tax))
}
