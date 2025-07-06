package main

import (
	"fmt"
	"time"
)

func main() {
	// Making a channel:
	// Note : channel cannot be used with Go routines
	greeting := make(chan any) // creating a channel and channel can be of any type
	greet := "Hello channel"
	items := []any{1, 2, 3, 4, 5, 7, 8, 9, "Numbers"}

	// Sender
	go func() { // Go Routine Anonymous function
		greeting <- greet // Channel receiving the greet value
		greeting <- items // Channel receiving the items slice
	}()

	// Receiver
	go func() {
		receiver := <-greeting // Channel sending the value to a variable
		fmt.Println(receiver)

		receiver = <-greeting // Channel sending the value to a variable
		fmt.Println(receiver)
	}()

	fmt.Println(greeting)

	time.Sleep(2 * time.Second)
}
