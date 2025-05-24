package main

import "fmt"

// func swap[T any](a, b T) (T, T) { // generic function
// 	return b, a
// }

type Stack[T any] struct { // generic struct
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) PrintStack() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack Elements : ")
	for _, v := range s.elements {
		fmt.Println(v)
	}

}

func main() {

	// x, y := 1, 2
	// x, y = swap(x, y)

	// fmt.Println(x, y)

	// a, b := "John", "Jane"
	// a, b = swap(a, b)
	// fmt.Println(a, b)

	// intStack := Stack[int]{}
	// intStack.Push(1)
	// intStack.Push(2)
	// intStack.Push(3)
	// intStack.Push(4)
	// intStack.Push(5)

	// fmt.Println(intStack.Pop())
	// fmt.Println(intStack.Pop())
	// fmt.Println(intStack.Pop())
	// fmt.Println(intStack.Pop())
	// fmt.Println(intStack.Pop())
	// fmt.Println(intStack.isEmpty())
	// intStack.PrintStack()

	stringStack := Stack[string]{}
	stringStack.Push("Jane")
	stringStack.Push("Jack")
	stringStack.Push("John")
	stringStack.Push("Jones")

	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.isEmpty())
	stringStack.PrintStack()

	// custome type generics :
	var age any

	age = 32
	age = "32"

	fmt.Println(age)

}
