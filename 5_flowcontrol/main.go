package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 175.43
	fmt.Println("Price", strconv.FormatFloat(kayakPrice, 'f', 2, 64))

	fmt.Println("=========================")
	fmt.Println("if")
	if kayakPrice > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice < 100 {
		fmt.Println("Price is less than 100")
	} else if kayakPrice > 200 && kayakPrice < 300 {
		fmt.Println("Price is between 200 and 300")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}

	fmt.Println("=========================")
	fmt.Println("Initialization Statement with an if")

	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("=========================")
	fmt.Println("for Loops")
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	for counter = 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
		// counter++
	}

	fmt.Println("=========================")
	fmt.Println("switch Statements")
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough // same as (case 'K', 'k':)
		case 'k':
			fmt.Println("K or k at position", index)
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)

		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}

	fmt.Println("=========================")
	fmt.Println("switch with initialization statement")
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}

	fmt.Println("=========================")
	fmt.Println("Label Statements (goto)")
	counter1 := 0
target:
	fmt.Println("Counter", counter1)
	counter1++
	if counter1 < 5 {
		goto target
	}
}
