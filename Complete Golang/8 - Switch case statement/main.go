package main

import (
	"fmt"
	"time"
)

func main() {
	// Single switch case :
	// var key string
	// fmt.Println("enter the key : ")
	// fmt.Scanf("%s", &key)

	// switch key {
	// case "w":
	// 	fmt.Println("Moving forward")
	// case "s":
	// 	fmt.Println("Moving backwards")
	// case "a":
	// 	fmt.Println("Moving leftward")
	// case "d":
	// 	fmt.Println("Moving rightward")
	// default:
	// 	fmt.Println("Please enter correct keys to move")
	// }
	// no need enter break keyword

	// Multi switch case :
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its workday")
	}

}
