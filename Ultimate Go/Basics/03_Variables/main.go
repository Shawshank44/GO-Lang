package main

import "fmt"

var global_variable string = "Jack" // global variables

func main() {
	// integer variable :
	var age int = 20
	fmt.Println(age)

	//string variable :
	var name string = "John Doe"
	fmt.Println(name)

	//floating point variable :
	var num float32 = 37.56
	fmt.Println(num)

	//boolean variable :
	var istrue bool = true
	fmt.Println(istrue)

	// shorthand for variables :
	num1 := 34
	fmt.Println(num1)

	// Declare multiple variables :
	var w, x, y, z int = 1, 3, 5, 7
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Naming conventions:
	//Pascal cases (Structs, Interfaces, Enums) (mainly used!):
	var PascalCase string = "PascalCase"
	fmt.Println(PascalCase)

	// Snakecase (ex : user_id : separated by "_" underscore):
	var snake_case string = "snake_case"
	fmt.Println(snake_case)

	// Upper case (specially used in constants):
	const UPPERCASE string = "UPPERCASE"
	fmt.Println(UPPERCASE)

	// mixedCase :
	var mixedCase string = "userName"
	fmt.Println(mixedCase)

	//constants :
	const (
		PI        = 3.14
		PORT      = 8000
		SUPERUSER = "user.1"
	)
	fmt.Println(SUPERUSER)

}
