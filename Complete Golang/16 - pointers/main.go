package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", num)
}

func main() {
	num := 1

	changeNum(&num)
	fmt.Println("after changeNum in main", num)

	fmt.Println("Memory address", &num)

}
