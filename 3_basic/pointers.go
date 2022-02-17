package main

import (
	"fmt"
	"sort"
)

func main() {

	first := 100
	var second *int
	fmt.Println("Second:", second)
	third := 105
	second = &third

	fouth := &second
	fifth := &fouth
	first++
	*second++
	fmt.Println("First:", first)
	fmt.Println("Second:", *second)
	fmt.Println("Third:", third)
	fmt.Println("fouth:", fouth)
	fmt.Println("fouth val:", **fouth)
	fmt.Println("fifth:", *fifth)
	fmt.Println("fifth val:", ***fifth)

	names := []string{"Alice", "Charlie", "Bob"}
	// secondName := &names[1]
	charlieIndex := pointerOfElement("Charlie", names)
	fmt.Println(*charlieIndex)
	sort.Strings(names[:])
	fmt.Println(*charlieIndex)
}

func pointerOfElement(element string, data []string) *string {
	for _, v := range data {
		if element == v {
			return &v
		}
	}
	return nil //not found.
}
