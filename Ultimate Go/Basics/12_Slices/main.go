package main

import "fmt"

func main() {

	// declarations and syntax
	// var slc [] int // aka: nil slice
	// slc := [] int {}

	// declartion using make function
	// make([]datatype, length, capacity)
	slc := make([]int, 0, 10)
	// append function adds element to the slice
	slc = append(slc, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	slc = append(slc, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	// Join two slices with append :
	slc1 := []int{21, 22, 23, 24, 25}
	slc = append(slc, slc1...)

	// copy fuction
	shortSLC := slc[:len(slc)-10]
	cpySLC := make([]int, len(shortSLC))
	copy(cpySLC, shortSLC)

	//length and capacity funtion:
	// copy slice:
	// fmt.Println("Copy of slice is ", cpySLC)
	// core slice
	// fmt.Println(slc)
	// fmt.Println("The length of the slice is : ", len(slc))
	// fmt.Println("The capacity of the slice is : ", cap(slc))

	// slice operator :
	// fmt.Println("Slice operator : ", slc[0:5])

	// interations in Slice :
	// for i := 0; i < len(slc); i++ {
	// 	fmt.Println("Index : ", i, "Value : ", slc[i])
	// }

	// for i, v := range slc {
	// 	fmt.Printf("Index : %d ; Value : %d \n", i, v)
	// }

	// accessing and modifying the elements in a slice :
	// ele := []int{0, 18, 6, 10, 5}
	// ele[3] = 10 // modifying
	// fmt.Println(ele)
	// fmt.Println("Element at index 4 is ", ele[4])

	// slices package :
	// slices.Sort(ele) // sorts the elements orderly
	// fmt.Println(ele)
	// i, f := slices.BinarySearch(ele, 5) // searches the element and returns index and bool(true/false)
	// fmt.Println(i, f)

	// Multi dimentional slices :
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)

}
