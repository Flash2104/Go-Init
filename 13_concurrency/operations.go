package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	channel := make(chan float64, 2)
	ind := 0
	for category, group := range data {
		ind++
		go group.TotalPrice(category, ind*1000, channel)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("-- Starting to receive from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println(len(channel), cap(channel))
		fmt.Println("-- channel read pending",
			len(channel), "items in buffer, size", cap(channel))
		value := <-channel
		fmt.Println("-- channel read complete", len(channel), "items in buffer, size", cap(channel))
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, wait int, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		// fmt.Println(category, "product:", p.Name)
		total += p.Price
		// time.Sleep(time.Millisecond * time.Duration(wait))
	}

	fmt.Println(category, "channel sending", ToCurrency(total))
	fmt.Println(category, "subtotal:", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}
