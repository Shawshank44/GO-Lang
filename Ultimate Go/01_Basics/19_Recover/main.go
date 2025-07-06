package main

import "fmt"

func main() {
	Process(-1)
	fmt.Println("Returned from process")
}

func Process(input int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered : ", r)
		}
	}()
	if input < 0 {
		panic("input must be a non-negative number")
	} else {
		fmt.Println("Enter number is : ", input)
	}
}

/*
1. The `main()` function starts and calls `Process(-1)`.
2. Inside `Process()`, a deferred function is defined to handle any panic using `recover()`.
3. The condition `if input < 0` is true since input is `-1`.
4. A `panic` is raised with the message `"input must be a non-negative number"`.
5. Normal execution stops, and control moves to the deferred function.
6. The deferred function catches the panic using `recover()` and prints `"Recovered : input must be a non-negative number"`.
7. `Process()` exits gracefully after recovering.
8. Control returns to `main()`, and `"Returned from process"` is printed.
*/
