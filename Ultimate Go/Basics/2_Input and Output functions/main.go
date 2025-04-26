package main

import "fmt"

func main() {
	// Output functions :
	// print
	// var i, j string = "Hello", "World"
	// fmt.Print(i)
	// fmt.Print(j)

	// New line
	// fmt.Print(i, "\n")
	// fmt.Print(j, "\n")

	//Println (one line)
	// fmt.Println(i, j)

	//printf :
	/*
		%v is used to print the value of the arguments
		%T is used to print the type of the arguments
	*/
	// fmt.Printf("The value is %v and the type is %T \n", i, i)
	// fmt.Printf("The value is %v and the type is %T", j, j)

	// Input functions :
	// Scan
	// var a, b int
	// fmt.Print("Enter two numbers: ")
	// fmt.Scan(&a, &b) // User inputs values separated by space
	// fmt.Println("You entered:", a, "and", b)

	// var x, y string
	// fmt.Print("Enter two words : ")
	// fmt.Scan(&x, &y)
	// fmt.Println("You entered:", x, "and", y)

	//Scanln
	// var name string
	// fmt.Print("Enter your full name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Hello,", name)

	// //Scanf
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Hello,", name, "Age:", age)

}
