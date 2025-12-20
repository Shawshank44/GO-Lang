package main

import (
	"fmt"
	"strconv"
)

func main() {
	// converting string(type) number to integer(type) number
	numStr := "12345"

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num + 1)

	// parsing as per bit size in integer
	parint, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parint)

	// Parsing float
	floatstr := "3.14"
	fval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed float : %.2f \n", fval)

	// binary to decimal
	binarystr := "1010"
	decimal, err := strconv.ParseInt(binarystr, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decimal)

	// Hexa decimal to integer
	Hexstr := "FF"
	Hex, err := strconv.ParseInt(Hexstr, 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(Hex)
}
