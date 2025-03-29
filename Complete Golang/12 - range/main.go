package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	sum := 0

	for _, num := range nums { // _ is index
		sum = sum + num
	}

	fmt.Println(sum)

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// returns unicode of a string
	for i, c := range "GOLANG" {
		fmt.Println(i, c)
	}
}
