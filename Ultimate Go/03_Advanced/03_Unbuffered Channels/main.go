package main

import (
	"fmt"
)

func main() {
	//By default channel is unbuffered
	/*
		Unbuffered Channels
			Definition: Channels with no capacity (make(chan int)) — data is passed only when both sender and receiver are ready.

			Behavior:

			A send operation (ch <- value) will block until another goroutine is ready to receive.

			A receive operation (val := <-ch) will block until another goroutine sends a value.

			Use Case: Useful for synchronization between goroutines.
	*/
	ch := make(chan int)

	go func() {
		ch <- 300
	}()

	receiver := <-ch
	fmt.Println(receiver)

}
