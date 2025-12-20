package main

import (
	"flag"
	"fmt"
)

func main() {
	// fmt.Println("command : ", os.Args[0])

	// fmt.Println("Arguement1 : ", os.Args[1])
	// for i, args := range os.Args {
	// 	fmt.Println("Arguement ", i, " : ", args)
	// }

	// go run main.go Hello World (command example)

	// define flags:
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "userX", "Name of the user")
	flag.IntVar(&age, "age", 1, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Gender : ", male)

	// go run main.go -name "James ding" -age 35  (Command example)
	// go run main.go --help (command to know how the flags are programmed)
}
