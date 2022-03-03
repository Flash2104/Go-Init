package main

import (
	"fmt"
)

func main() {
	fmt.Println("main func started")
	// timer := time.NewTimer(0)
	// a := <-timer.C
	// fmt.Println("Start ticker:", a.Format(time.Nanosecond.String()))

	CalcStoreTotal(Products)
	// fmt.Println("End ticker:", (<-timer.C).Format(time.Nanosecond.String()))
	// timer.Stop()
	fmt.Println("main func complete")
}
