package main

import (
	"fmt"
	"time"
)

func main() {
	// done := make(chan int)
	// go func() {
	// 	fmt.Println("Working")
	// 	time.Sleep(2 * time.Second)
	// 	done <- 0
	// }()
	// <-done
	// fmt.Println("Finished")

	// ch := make(chan int)
	// go func() {
	// 	fmt.Println("Sending.....")
	// 	ch <- 9 // blocking until the value is received
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Sent value")
	// }()
	// value := <-ch // Blocking until a value is sent
	// time.Sleep(3 * time.Second)
	// fmt.Println(value)

	// Synchronization of multiple channels
	// numGORoutines := 3
	// done := make(chan int, 3)

	// for i := range numGORoutines {
	// 	go func(id int) {
	// 		fmt.Printf("GoRoutines %d working...\n", id)
	// 		time.Sleep(time.Second)
	// 		done <- id
	// 	}(i)
	// }

	// for range numGORoutines {
	// 	<-done // wait for each Go routines to finish
	// }
	// fmt.Println("All Go routines are finished")

	// Synchronizing channels for data exchange :
	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- fmt.Sprint("Hello ", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data) // Closes the
	}()

	for value := range data {
		fmt.Println("Recevied value : ", value, " : ", time.Now())
	}
}
