package main

import "fmt"

func main() {
	// fmt.Println(Factorial(5))
	// fmt.Println(Factorial(10))

	// fmt.Println(SumOfDigits(9))
	// fmt.Println(SumOfDigits(12))
	// fmt.Println(SumOfDigits(12345))

	n := 12
	fmt.Printf("Fibonacci(%d) = %d\n", n, Fibonacci(n))
}

func Factorial(n int) int {
	// Base case : factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case : factorial of n is n * factorial (n - 1)
	return n * Factorial(n-1)
}

func SumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + SumOfDigits(n/10)
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
	Uses of Recursion :
	Mathematical algorithms
	tree and graph traversal
	divide and conquer algorithms
*/
