package main

import "fmt"

func main() {
	// var ptr *int
	// var a int = 10
	// ptr = &a // referencing the pointer
	// fmt.Println(a)
	// fmt.Println(ptr)
	// fmt.Println(*ptr) // dereferencing the pointer

	v := 8
	fmt.Println(ModifyValue(&v))
}

func ModifyValue(ptr *int) int {
	*ptr++
	return *ptr
}
