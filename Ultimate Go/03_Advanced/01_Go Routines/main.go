package main

import (
	"fmt"
	"time"
)

func SayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello From Go routine")
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printAlphabets() {
	for _, letter := range "ABCDE" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func DoWork() error {
	// Stimulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("error occured in DOWORK")
}

func main() {
	var err error
	fmt.Println("staring the main function")
	go SayHello() // Go Routine
	go printNumbers()
	go printAlphabets()

	go func() { // we cannot use 'go' keyword in variable so we use anonymous function
		err = DoWork()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Work Done")
		}
	}()

	fmt.Println("after say hello function")

	time.Sleep(2 * time.Second)

}
