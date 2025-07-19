package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. Non Blocking Receive operation
	// ch := make(chan int)

	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received : ", msg)
	// default:
	// 	fmt.Println("No messages available.")
	// }

	// // 2. Non Blocking send operation
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent Message")
	// default:
	// 	fmt.Println("Channel is not ready to receive.")
	// }

	// 3. Non Blocking Operation in Real-time system
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received : ", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true

}
