package main

import "fmt"

func main() {
	Process()
	fmt.Println("Returned from process")
}

func Process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered : ", r)
		}
	}()
	fmt.Println("Start process")
	panic("Something went wrong!")
	// fmt.Println("End process") // will not be executed
}
