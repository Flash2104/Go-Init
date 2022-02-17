package main

import "fmt"

func main() {
	printHello()
	for i := 0; i < 5; i++ {
		i = i
		printNumber(i)
	}
	fmt.Scanln()
}

func printHello() {
	fmt.Println("Hello, Go")
}

func printNumber(number int) {
	fmt.Println(number)
}
