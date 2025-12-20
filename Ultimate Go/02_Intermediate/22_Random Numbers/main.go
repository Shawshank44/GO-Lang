package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(10))     // Generating a random number (0 to 10)
	fmt.Println(rand.Intn(10) + 1) // Adding lower limit shift from (1 to 10)
	fmt.Println(rand.Int())        // generate random digits

	// Seeding random numbers :
	// val := rand.New(rand.NewSource(40))
	val := rand.New(rand.NewSource(time.Now().Unix())) // Using time package
	fmt.Println(val.Intn(101))

	// Generating float numbers :
	fmt.Println("Random float numbers", rand.Float64())

	// Rolls N Dice:
	for {
		// Menu:
		fmt.Println("Welcome to the Dice")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. exit")
		fmt.Print("Enter your choice : ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, Please enter 1 or 2")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for nothing :> bye!")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Println("Total : ", die1+die2)

		// Ask if the user wants to roll again
		fmt.Print("Do you want to roll again (y/n): ")
		var rollagain string
		_, err = fmt.Scan(&rollagain)
		if err != nil || (rollagain != "y" && rollagain != "n") {
			fmt.Println("Invalid input ")
			break
		}
		if rollagain == "n" {
			fmt.Println("Thankyou for playing! Good bye")
			break
		}
	}
}
