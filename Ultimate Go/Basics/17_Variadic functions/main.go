package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}

	Msg, Total := Sum("The Sum of nums is : ", nums...)

	fmt.Println(Msg, Total)
}

// Note : Variadic ...params must be last ...params always!!!!

func Sum(msg string, nums ...int) (string, int) { // ... is denoted as Ellipsis
	total := 0
	for _, v := range nums {
		total += v
	}
	return msg, total
}
