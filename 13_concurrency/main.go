package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main func started")
	CalcStoreTotal(Products)
	time.Sleep(time.Second * 5)
	fmt.Println("main func complete")
}
