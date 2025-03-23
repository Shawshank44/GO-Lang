package main

import "fmt"

func main() {

	var nums [4]int // declaration

	fmt.Println(len(nums)) // length of an array

	// initialization
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(nums)

	integers := [...]int{1, 2, 3, 4, 5, 6} // initialization and declaration
	fmt.Println(integers)

	Strings := [...]string{"John", "Doe", "alice", "sharpe"}
	fmt.Println(Strings)

	//initialization on specific index :
	arr1 := [...]int{1: 20, 4: 30, 3: 50}
	fmt.Println(arr1)

	// printing array using for loop :
	for i := 0; i < len(integers); i++ {
		fmt.Println(integers[i])
	}

}
