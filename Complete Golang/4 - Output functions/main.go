package main

import "fmt"

func main() {
	// print
	var i, j string = "Hello", "World"
	// fmt.Print(i)
	// fmt.Print(j)

	// New line
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	//Println (one line)
	fmt.Println(i, j)

	//printf :
	/*
		%v is used to print the value of the arguments
		%T is used to print the type of the arguments
	*/
	fmt.Printf("The value is %v and the type is %T \n", i, i)
	fmt.Printf("The value is %v and the type is %T", j, j)

}
