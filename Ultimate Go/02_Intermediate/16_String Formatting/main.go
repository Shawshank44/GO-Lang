package main

import "fmt"

func main() {

	// Number alignment
	Num := 42
	fmt.Printf("%05d\n", Num)

	// String Alignment:
	Message := "Hello"
	fmt.Printf("|%10s|\n", Message)
	fmt.Printf("|%-10s|\n", Message)

	//String Interpolation :
	message1 := "Hello \n world!"
	message2 := `Hello \n World!`
	fmt.Println(message1)
	fmt.Println(message2)
}
