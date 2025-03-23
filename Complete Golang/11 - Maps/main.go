package main

import (
	"fmt"
)

func main() {
	// creating a map :
	age := make(map[string]int)

	// setting an element :
	age["alice"] = 30
	age["john"] = 40
	age["Doe"] = 25
	age["NUN"] = 3656

	fmt.Println(age)
	fmt.Println(age["alice"])

	fmt.Println(len(age)) // length of the map
	//updating the element
	age["alice"] = 28

	//deleting the element
	delete(age, "NUN")

	clear(age) //clears all the functions in the map

	fmt.Println(age)

	// other ways to create map
	ageing := map[string]int{
		"Jack":      20,
		"jasmine":   30,
		"kevin":     25,
		"alejendra": 35,
		"sean":      33,
	}
	fmt.Println(ageing)

	// Iterate Over Maps :
	for k, v := range ageing {
		fmt.Printf("%v : %v, ", k, v)
	}
}
