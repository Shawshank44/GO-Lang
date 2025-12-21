package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic Buffer channel example :
	// ch := make(chan any, 2)

	// ch <- 10
	// ch <- "Hello Channel"
	// // ch <- 300 // Blocks because channel capacity is only 2
	// go func() { // will not block if passed inside Go routine
	// 	ch <- 300
	// }()

	// // receiver1 := <-ch
	// // receiver2 := <-ch
	// // receiver3 := <-ch // blocks cannot receive
	// fmt.Println(<-ch, <-ch, <-ch)

	// Scenrio Blocking on send only if the buffer is full :
	ch := make(chan any, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Go routine 2 sec timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Recieved : ", <-ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3 // Blocks because the buffer is full
	fmt.Println("Blocking ends")
	fmt.Println("Received : ", <-ch)
	fmt.Println("Received : ", <-ch)

	// Scenrio Blocking on receive only if the buffer is empty :
	// CH := make(chan int, 2)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	CH <- 1
	// 	CH <- 2
	// }()
	// fmt.Println("Value 1: ", <-CH)
	// fmt.Println("Value 2: ", <-CH)
	// fmt.Println("End of program")
}

/*
	🔹 Buffered Channels
	Definition: Channels with a defined capacity (make(chan int, 3)) — data can be sent without immediate receiving, up to the buffer limit.

	Behavior:

	A send operation only blocks if the channel is full.

	A receive operation only blocks if the channel is empty.

	Use Case: Useful when you want to allow some asynchronous communication between goroutines.

*/
