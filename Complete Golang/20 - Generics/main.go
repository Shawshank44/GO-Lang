package main

import "fmt"

func PrintSlice[T int | string](items ...T) { // we can use Interface or any and also add limiters
	for _, v := range items {
		fmt.Println(v)
	}
}

type Stack[T interface{}] struct { // we can comparable, any and interface.
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	names := []string{"John", "Jack", "carey"}
	PrintSlice(nums...)
	PrintSlice(names...)

	myStack := Stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(myStack)

}
