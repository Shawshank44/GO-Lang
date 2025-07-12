package main

import "fmt"

//Send only channel:
func SendOnlyChannel(ch chan<- any) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func ReceiveOnlyChannel(ch <-chan any) { // type of channel must be same even used "any or interface{}"
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	ch := make(chan any) // Bidirectinal channel which can accept both send and receive

	//Send only channel
	SendOnlyChannel(ch)

	// Receive only channel
	ReceiveOnlyChannel(ch)
}
