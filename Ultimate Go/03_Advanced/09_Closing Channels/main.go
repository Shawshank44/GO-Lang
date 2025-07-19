package main

import "fmt"

func main() {
	// Simple Closing channel :
	// ch := make(chan int)

	// go func() {
	// 	for i := range 5 {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// Receiving from a closed channel :
	// ch := make(chan int)
	// close(ch)
	// val, ok := <-ch
	// if !ok {
	// 	fmt.Println("Channel is closed")
	// 	return
	// }
	// fmt.Println(val)

	// Range from a closed channel :
	// Note : Always close channels from the sending point do not close from receiving point
	// ch := make(chan int)
	// go func() {
	// 	for i := range 5 {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// Run-time Panic while closing channels multiple times (*leads to error) :
	// ch := make(chan int)
	// go func() {
	// 	close(ch)
	// 	close(ch)
	// }()
	// time.Sleep(time.Second)

	// producer and filter :
	ch1 := make(chan int)
	ch2 := make(chan int)

	producer := func(ch chan<- int) {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}

	filter := func(in <-chan int, out chan<- int) {
		for val := range in {
			if val%2 == 0 {
				out <- val
			}
		}
		close(out)
	}

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}
