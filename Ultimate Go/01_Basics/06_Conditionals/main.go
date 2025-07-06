package main

import "fmt"

func main() {

	age := 101

	// basic if else :
	// if age >= 18 {
	// 	fmt.Print("You can drive the vehicles")
	// } else {
	// 	fmt.Print("You are not eligible to drive")
	// }

	// else if :
	// var time int
	// fmt.Print("Enter the time as per 24hrs format: ")
	// fmt.Scan(&time)
	// if time <= 11 {
	// 	fmt.Println("Good morning.")
	// } else if time <= 15 {
	// 	fmt.Println("Good afternoon.")
	// } else {
	// 	fmt.Println("Good evening.")
	// }

	// if else ladder :
	if age <= 0 || age > 100 {
		fmt.Println("Invalid input must be greater than 0 or less than 100")
	} else if age >= 18 && age < 21 {
		fmt.Println("You can only drive two wheeler")
	} else if age >= 21 {
		fmt.Println("You can drive both two wheeler and four wheeler and above")
	} else {
		fmt.Println("You can only drive cycle")
	}

}
