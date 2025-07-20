package main

import (
	"fmt"
	"time"
)

/*
	In Go (Golang), the time package provides a Timer type that allows you to execute code after a specified duration. It's very useful for delaying execution or implementing timeouts.
*/

func longRunningOperations() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// Basic Timer use
	// fmt.Println("Starting app.")
	// timer := time.NewTimer(2 * time.Second)
	// fmt.Println("Waiting for timer")
	// stopped := timer.Stop() // stopping the time
	// if stopped {
	// 	fmt.Println("timer stopped")
	// }
	// fmt.Println("timer Reset")
	// timer.Reset(time.Second)
	// <-timer.C
	// fmt.Println("Timer expired")

	// Timeout use
	// timeout := time.After(3 * time.Second)
	// done := make(chan bool)

	// go func() {
	// 	longRunningOperations()
	// 	done <- true
	// }()

	// select {
	// case <-timeout:
	// 	fmt.Println("Operation times out")
	// case <-done:
	// 	fmt.Println("Operation completed")
	// }

	// Scheduling delayed use :
	// timer := time.NewTimer(2 * time.Second)
	// go func() {
	// 	<-timer.C
	// 	fmt.Println("Delayed operation executed")
	// }()
	// fmt.Println("Waiting.....")
	// time.Sleep(3 * time.Second)
	// fmt.Println("End of program")

	// Multiple timers
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	for range 2 {
		select {
		case <-timer1.C:
			fmt.Println("Timer1 expired")
		case <-timer2.C:
			fmt.Println("Timer2 expired")
		}
	}

}
