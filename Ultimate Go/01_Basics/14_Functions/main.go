package main

import "fmt"

func main() {
	// fmt.Println(Sum(10, 5))

	//Anonymous function aka (closures):
	// func(a, b int) {
	// 	fmt.Println(a + b)
	// }(20, 3)

	// assigning function to a variable
	// multiply := func(a, b int) int {
	// 	return a * b
	// }

	// fmt.Println(multiply(10, 10))

	// passing function as an arguement
	result := Operation(5, 3, Sum)
	fmt.Println(result)

	//Returning and using a function
	mult := CreateMultiplyer(2)
	fmt.Println(mult(6))

}

func Sum(a, b int) int {
	return a + b
}

// function that takes function as an arguements:
func Operation(x, y int, operate func(int, int) int) int {
	return operate(x, y)
}

// functions that returns a function :
func CreateMultiplyer(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
