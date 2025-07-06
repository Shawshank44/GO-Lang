package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"

	fmt.Println(str)
	fmt.Println("The length of the string is : ", len(str))

	// concatination of string
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])         // to get ASCII value of a particular alphabet
	fmt.Println(string(str[0])) // to get particular alphabet
	fmt.Println(str[0:5])       // Indexing of a string

	// Standard library functions for string :
	// 1. integer to string
	num := 18
	stri := strconv.Itoa(num)
	fmt.Println("Length after converting number into string", len(stri))

	//2. String Splitting :
	fruits := "apple-orange-banana"
	// comma := strings.Split(fruits, ",")
	// fmt.Println(comma)
	dash := strings.Split(fruits, "-")
	fmt.Println(dash)

	// 3. String join (concatination)
	countries := []string{"India", "Germany", "Italy", "France"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	// 4. strings contains?
	strcon := "Hello Hi"
	fmt.Println(strings.Contains(strcon, "Hello"))

	//5. Strings replace :
	strrep := "Hello Replace"
	replaced := strings.Replace(strrep, "Replace", "World", 1)
	fmt.Println("Before replacing : ", strrep)
	fmt.Println("After replacing : ", replaced)

	//6. Trimming strings (spaces) :
	strspc := " Hello Everyone!	"
	fmt.Println(strspc)
	fmt.Println(strings.TrimSpace(strspc))

	//7. Uppercase and lowercase :
	strul := "UppErcAsE aNd lOwErCaSe"
	fmt.Println("converting string to lower case : ", strings.ToLower(strul))
	fmt.Println("converting string to UPPER CASE : ", strings.ToUpper(strul))

	//8. Repeating strings
	fmt.Println(strings.Repeat(" Foo ", 3))

	//9. Occurance and counting
	fmt.Println(strings.Count("NANNANA", "A"))

	//10. Prefix and suffix in string:
	fmt.Println(strings.HasPrefix("Hello", "H"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	// REGEX(regular expressions) :
	strex := "1Hello, 123 Go! 11"
	rex := regexp.MustCompile(`\d+`)
	fmt.Println(rex.FindAllString(strex, -1))

	// Unicode package :
	struni := "Hello おはよう"
	fmt.Println(utf8.RuneCountInString(struni))

	// String Builder :
	var Builder strings.Builder

	// (1) Writing some strings :
	Builder.WriteString("Hello")
	Builder.WriteString(", ")
	Builder.WriteString("world!")

	// (2) Convert builder to a string
	res := Builder.String()
	fmt.Println(res)

	//(3) Using Writerune to add a character :
	Builder.WriteRune(' ')
	Builder.WriteString("How are you")
	res = Builder.String()
	fmt.Println(res)

	//(4) to write new string :
	Builder.Reset()
	Builder.WriteString("Starting new!!")
	res = Builder.String()
	fmt.Println(res)

}
