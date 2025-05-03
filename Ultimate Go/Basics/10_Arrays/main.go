package main

import "fmt"

func main() {
	// var arr [10]int

	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5
	// arr[8] = 10
	// fmt.Println(arr)

	// names := [5]string{"John", "Jack", "Windy", "Bob", "Ceaser"}
	// fmt.Println(names)

	// fruits := [5]string{1: "apple", 2: "Mango", 4: "orange", 3: "Grapes"} // assigning elements to a specific index
	// fmt.Println(fruits)

	// elements := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fmt.Println(len(elements)) // to know the size of the array

	// for i := 0; i < len(elements); i++ { // normal interation
	// 	fmt.Println("Elements at the index,", i, " : ", elements[i])
	// }

	// for _, v := range elements { // range based iteration
	// 	fmt.Printf("Value : %d \n", v)
	// }

	// "_"(Underscore) is a blank indentifier if we do not want use a value we can assign underscore

	// comparing the arrays:
	// arr1 := [3]int{1, 2, 3}
	// arr2 := [3]int{10, 2, 3}

	// fmt.Println(arr1 == arr2)

	//Multi dimentional array :
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

}
