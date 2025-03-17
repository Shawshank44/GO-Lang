package main

import "fmt"

func main() {
	// for -> only construct in GO for looping

	//while loop in for
	// i := 1
	// for i <= 10 {
	// 	println(i)
	// 	i++
	// }

	// infinite loop :
	// for {
	// 	println(i)
	// 	i++
	// }

	// for loop :
	// for i := 0; i < 10; i++ {
	// 	println(i)
	// }

	// // for range :
	for i := range 10 {
		fmt.Println(i)
	}
}
