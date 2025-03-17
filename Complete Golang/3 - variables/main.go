package main

import "fmt"

const PI float32 = 3.14 // constant varaibles cannot be changed

func main() {

	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// Declare multiple variables :
	var w, x, y, z int = 1, 3, 5, 7

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//Strings:
	// var username string = "Golang"
	var username = "Golang"
	fmt.Println(username)

	//Boolean :
	// var is_true bool = true
	var is_true = true
	fmt.Println(is_true)

	//Integers :
	// var age int = 34s
	var age = 34
	fmt.Println(age)

	// short hand
	name := "GOLANG"
	fmt.Println(name)

	// Declaration and initialization
	var Greet string

	Greet = "Hello Good morning"

	println(Greet)

	// floats :
	var kilo float32 = 60.23
	var gram float64 = 60.45
	println(kilo)
	println(gram)

	fmt.Printf("The value of Constant PI is %v", PI)

	// Rules of variables :
	// A variable name must start with a letter or an underscore character (_)
	// A variable name cannot start with a digit
	// A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
	// Variable names are case-sensitive (age, Age and AGE are three different variables)
	// There is no limit on the length of the variable name
	// A variable name cannot contain spaces
	// The variable name cannot be any Go keywords

}
