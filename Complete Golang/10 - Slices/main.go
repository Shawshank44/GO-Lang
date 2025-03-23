package main

import (
	"fmt"
	"slices"
)

func main() {

	// var num []int // declartion
	// fmt.Println(num)
	// fmt.Println(num == nil)
	//int, minimum size and minimum capacity
	// var nums = make([]int, 0, 5) // declaration
	// nums := []int{}
	// fmt.Printf("The number of minimum size nums can store is : %v \n", len(nums))
	// fmt.Printf("The number of minimum capacity nums can store is : %v \n", cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// fmt.Println(nums)
	// fmt.Printf("The number of minimum size nums can store is : %v \n", len(nums))
	// fmt.Printf("The number of minimum capacity nums can store is : %v \n", cap(nums))

	//Append One Slice To Another Slice :
	// numbers := []int{1, 2, 3, 4, 6}
	// numbers2 := []int{7, 8, 9, 10}
	// fmt.Println(numbers, numbers2)
	// numbersjoin := append(numbers, numbers2...)
	// fmt.Println(numbersjoin)

	// Copy method :
	// slicers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// // Original slice
	// fmt.Printf("numbers = %v\n", slicers)
	// fmt.Printf("length = %d\n", len(slicers))
	// fmt.Printf("capacity = %d\n", cap(slicers))
	// //copy slice
	// shortslice := slicers[:len(slicers)-10]
	// numbersCopy := make([]int, len(shortslice))
	// copy(numbersCopy, shortslice)
	// fmt.Printf("numbersCopy = %v\n", numbersCopy)
	// fmt.Printf("length = %d\n", len(numbersCopy))
	// fmt.Printf("capacity = %d\n", cap(numbersCopy))

	// slice operator :
	// fmt.Println(slicers[0:5]) // from 0th index to 5th index(which exclude the 5th element)

	// slice package :
	// fmt.Println(slices.Contains(slicers, 10)) // checks whether the element is present or not (returns boolean)
	// fmt.Println(slices.Index(slicers, 3))     // returns the element present in the index
	// slices.Sort(slicers)                      // sorts the element in an order
	// fmt.Println(slicers)
	// slices.Reverse(slicers)
	// fmt.Println(slicers)

	//Binary search (Works only on sorted slices → Undefined behavior if the slice is unsorted.):
	nums := []int{2, 4, 6, 8, 10, 12}
	// Searching for an existing value
	index, found := slices.BinarySearch(nums, 8)
	fmt.Println("Index:", index, "Found:", found) // Output: Index: 3 Found: true

}
