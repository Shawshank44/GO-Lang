package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(operator("The Sum of two numbers is :", 13, 6))

	add, sub, mult := Sum_ops(10, 20)

	fmt.Println("By adding : ", add, "By Subtracting : ", sub, "By Multiplying : ", mult)

	// res, err := Compare(5, 2)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }
}

// Own Functions
// func operator(msg string, a, b int) (string, int) {
// 	return msg, a + b
// }

// func Sum_ops(a, b int) (int, int, int) {
// 	add := a + b
// 	sub := a - b
// 	mult := a * b
// 	return add, sub, mult
// }

func Sum_ops(a, b int) (add int, sub int, mult int) { // named return values
	add = a + b
	sub = a - b
	mult = a * b
	return // after naming the return values we can only just mention return.
}

func Compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil // when we do not have error to send here just type nil
	} else if b > a {
		return "b is greater than a", nil // error is an interface zero value of interface is nil
	} else {
		return " ", errors.New("Invalid statements") // for strings zero value is " "(blank string)
	}

}
