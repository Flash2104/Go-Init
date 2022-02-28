package data

import "fmt"

func init() {
	fmt.Println("Data init func invoked")
}

func GetData() []string {
	return []string{
		"Kayak",
		"Lifejacket",
		"Paddle",
		"Soccer Ball",
	}
}
