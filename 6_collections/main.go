package main

import "fmt"

func main() {
	fmt.Println("===================")
	fmt.Println("Arrays")
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
	names1 := [...]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names1)
	names2 := &names
	names[0] = "123s"
	fmt.Println(names2)

	same := names == names1
	fmt.Println("comparison:", same)

	for index, value := range names2 {
		fmt.Println("Index:", index, "Value:", &value)
	}

	fmt.Println("===================")
	fmt.Println("Slices")
	namesSlice := make([]string, 3, 6)
	namesSlice[0] = "Kayak"
	namesSlice[1] = "Lifejacket"
	namesSlice[2] = "Paddle"
	fmt.Println("namesSlice:", namesSlice)
	fmt.Println("len:", len(namesSlice))
	fmt.Println("cap:", cap(namesSlice))
	appendedNames := append(namesSlice, "Hat", "Gloves")
	namesSlice[0] = "Canoe"
	fmt.Println("names:", namesSlice)
	fmt.Println("appendedNames:", appendedNames)

}
