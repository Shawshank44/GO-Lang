package main

import "fmt"

func main() {

	// Printing functions:
	// fmt.Print("Hello")
	// fmt.Print("Print!")
	// fmt.Print(12, 456)

	// fmt.Println("Hello")
	// fmt.Println("Println!")
	// fmt.Println(120, 4580)

	// name := "John"
	// age := 25
	// fmt.Printf("Name : %s, Age : %d", name, age)

	// Formatting functions
	// sp := fmt.Sprint("Hello ", " World!", 123, 456)
	// fmt.Println(sp)
	// spln := fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Println(spln)
	// spf := fmt.Sprintf("Name : %s, Age : %d", "john", 34)
	// fmt.Println(spf)

	//Scanning functions :
	// var name string
	// var age int
	// fmt.Print("Enter your name and age : ")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// // fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name : %s, Age : %d \n", name, age)

	//Error Formatting :
	err := CheckAge(19)
	if err != nil {
		fmt.Println("Error : ", err)
	}

}

func CheckAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
