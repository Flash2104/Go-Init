package main

import (
	"fmt"
)

func main() {
	fmt.Println("===================")
	fmt.Println("Arrays")
	// var names [3]string
	// names[0] = "Kayak"
	// names[1] = "Lifejacket"
	// names[2] = "Paddle"
	// fmt.Println(names)
	// names1 := [...]string{"Kayak", "Lifejacket", "Paddle"}
	// fmt.Println(names1)
	// names2 := &names
	// names[0] = "123s"
	// fmt.Println(names2)

	// same := names == names1
	// fmt.Println("comparison:", same)

	// for index, value := range names2 {
	// 	fmt.Println("Index:", index, "Value:", &value)
	// }

	// fmt.Println("===================")
	// fmt.Println("Slices")
	// namesSlice := make([]string, 3, 6)
	// namesSlice[0] = "Kayak"
	// namesSlice[1] = "Lifejacket"
	// namesSlice[2] = "Paddle"
	// fmt.Println("namesSlice:", namesSlice)
	// fmt.Println("len:", len(namesSlice))
	// fmt.Println("cap:", cap(namesSlice))
	// appendedNames := append(namesSlice, "Hat", "Gloves")
	// namesSlice[0] = "Canoe"
	// fmt.Println("names:", namesSlice)
	// fmt.Println("appendedNames:", appendedNames)

	// moreNames := []string{"Hat Gloves"}
	// appendedNames = append(namesSlice, moreNames...)
	// fmt.Println("appendedNames:", appendedNames)

	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// products[0] = "Kayak"
	// products[1] = "Lifejacket"
	// products[2] = "Paddle"
	// products[3] = "Hat"
	allNames := products[0:]
	someNames := allNames[1:]
	allNames = append(allNames, "Gloves")
	products[2] = "Canoe"
	fmt.Printf("someNamesRef:%-p\n", &someNames)
	fmt.Println("someNames:", someNames)
	fmt.Printf("allNamesRef:%-p\n", &allNames)
	fmt.Println("allNames:", allNames)

	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// someNames := make([]string, 3)
	// copy(someNames, allNames)
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames", allNames)
	// products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// allNames := products[1:]
	// var someNames []string
	// copy(someNames, allNames)
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames", allNames)

	// fmt.Println("===================")
	// fmt.Println("Comparing Slices. Deep Equal")
	// p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	// p2 := make([]string, 4)
	// copy(p2, p1)
	// p2 = append(p2[:2], p2[3:]...)

	// fmt.Println("p2:", p2)
	// fmt.Println("Equal:", reflect.DeepEqual(p1, p2))

	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[2]string)(p1)
	array := *arrayPtr
	fmt.Println(array)
	fmt.Printf("Pointer: %-p", arrayPtr)
}
