package main

import "fmt"

func main() {
	Process(-10)
}

func Process(input int) {

	defer fmt.Println("Defered 1")
	defer fmt.Println("Defered 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After panic") // will not executed after panic
		// defer fmt.Println("Defered 3")// will not executed after panic
	}
	fmt.Println("Processing input : ", input)
}

/*

🧠 What Happens at Runtime?

main() calls Process(-10)

Inside Process:

Defered 1 and Defered 2 are scheduled.

Condition input < 0 is true:

"Before Panic" is printed.

Panic occurs.

As the stack unwinds:

Deferred functions are called in reverse order:

"Defered 2" is printed.

"Defered 1" is printed.

Program crashes with panic after running deferred functions.

*/
