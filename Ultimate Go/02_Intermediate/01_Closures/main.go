package main

import "fmt"

func main() {
	// sequence := Adder()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))
}

func Adder() func() int {
	i := 0
	fmt.Println("Previous value of i : ", i)

	return func() int {
		i++
		fmt.Println("added 1 to i ")
		return i
	}
}
