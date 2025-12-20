package main

import "fmt"

func main() {
	// message := "Hello\nGo"   // string with escape sequence character (\n to print on next line)
	// message1 := "Hello\tGo"  // string with escape sequence character (\t to print on space in between)
	// IsMessage := `Hello\nGo` // backticks (template literals) treats every this as string

	// fmt.Println(message)
	// fmt.Println(message1)
	// fmt.Println(IsMessage)

	// fmt.Println("Length of the message is : ", len(IsMessage)) // to find the length of the string
	// fmt.Println(message[0])                                    // returns the ASCII value
	// fmt.Println(string(message[0]))                            // to get the real value of the string

	// // String concatination :
	// greeting := "Hello "
	// name := " Alice"
	// fmt.Println(greeting + name)

	// // comparision (comparision is done on the basis of ASCII value):
	// fmt.Println(greeting == name) // returns bool (true/false)
	// fmt.Println(greeting > name)
	// fmt.Println(greeting < name)

	// Iteration of string :
	// Str := "This string is used for iteration"

	// normal iteration:
	// for i := 0; i < len(Str); i++ {
	// 	fmt.Println("Index : ", i, "Characters : ", string(Str[i]))
	// }

	// Range based iteration :
	// for i, chars := range Str {
	// 	fmt.Println("Index : ", i, "Characters : ", string(chars))
	// }

	// // counting of Utf8 characters in string
	// fmt.Println("Rune count : ", utf8.RuneCountInString(Str))

	// Runes :
	// var ch rune = 'a'

	// fmt.Println(ch)         // returns the ASCII value
	// fmt.Printf("%c \n", ch) // returns the actual value

	// // conversion of rune to string
	// fmt.Printf("type of ch is %T ", ch) // return int32 because rune's are AKA integers
	// ctstr := string(ch)
	// fmt.Println(ctstr)
	// fmt.Printf("type of ch is %T", ctstr)

	const WHATISUP = "どうした"
	fmt.Println(WHATISUP)

	for _, runeValue := range WHATISUP {
		fmt.Printf("%c\n", runeValue)
	}
}
