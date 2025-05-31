package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile a regex pattern
	regex := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	//test strings
	email1 := "user119@gmail.com"
	email2 := "invalid_mail"

	//Match
	fmt.Println("email1 : ", regex.MatchString(email1))
	fmt.Println("email2 : ", regex.MatchString(email2))

	// Capturing groups
	// Compile a regex pattern to capture date components
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date := "2024-07-30"

	// Find all Submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// source strings :
	str := "Hello World"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// Flags :
	// i - Case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	tes := "Golang is going great"

	fmt.Println("Match : ", re.MatchString(tes))
}
