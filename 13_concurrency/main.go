package main

import (
	"fmt"
)

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func main() {
	fmt.Println("main func started")
	// timer := time.NewTimer(0)
	// a := <-timer.C
	// fmt.Println("Start ticker:", a.Format(time.Nanosecond.String()))

	// CalcStoreTotal(Products)
	// // fmt.Println("End ticker:", (<-timer.C).Format(time.Nanosecond.String()))
	// // timer.Stop()
	// fmt.Println("main func complete")

	fmt.Println("----------------------------")
	fmt.Println("Sending and Receiving an Unknown Number of Values")

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		if details, open := <-dispatchChannel; open {
			fmt.Println("Dispatch To", details.Customer, ":", details.Quantity, "x", details.Name)
		} else {
			break
		}
	}

	// or Enumerating Channel Values
	fmt.Println("----------------------------")
	fmt.Println("Enumerating Channel Values")
	dispatchChannel = make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)

	for details := range dispatchChannel {
		fmt.Println("Dispatch To", details.Customer, ":", details.Quantity, "x", details.Name)
	}

	fmt.Println("----------------------------")
	fmt.Println("Restricting Channel Argument Direction")
	dispatchChannel = make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
}
