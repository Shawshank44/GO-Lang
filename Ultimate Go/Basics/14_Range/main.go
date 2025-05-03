package main

import "fmt"

func main() {

	// message := "Hello World"
	// for i, v := range message {
	// 	fmt.Println(i, v)
	// 	fmt.Printf("Index : %d, Rune : %c \n", i, v)
	// }

	// ages := map[string]int{
	// 	"Jack": 25,
	// 	"John": 30,
	// 	"Jeff": 56,
	// }

	// for k, v := range ages {
	// 	fmt.Println(k, v)
	// }

	slcs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slcs = append(slcs, 11, 12, 13, 14, 15)

	for i, v := range slcs {
		fmt.Println("Index : ", i, "value : ", v)
	}

}
