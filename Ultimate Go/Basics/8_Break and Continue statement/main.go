package main

import "fmt"

func main() {
	// fmt.Println("*****Break statement*****")
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// fmt.Println("*****continue statement*****")
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Both using together :
	for i := 1; i <= 10; i++ {
		if i == 8 {
			break // exits the loop when i is 8
		}

		if i%2 == 0 {
			fmt.Println(i, "is even")
			continue // skip to next iteration
		}

		fmt.Println(i, "is odd")
	}

}
