package main

import "fmt"

func main() {
	fmt.Println("Main function")
}

func init() {
	fmt.Println("Intialysing function 1")
}

func init() {
	fmt.Println("Intialysing function 2")
}

func init() {
	fmt.Println("Intialysing function 3")
}

/*
	In Go, the init function is a special function that is automatically executed before the main function, and it's used to perform setup tasks like initializing variables or configurations.

	Here are key points about the init function in single sentences:

	The init function is called automatically before the main function runs.

	It does not take any arguments and does not return anything.

	A package can have multiple init functions, even in different files.

	If multiple init functions exist, they are executed in the order of file compilation and appearance.

	It is mainly used for setup tasks like initializing state or checking preconditions.

*/
