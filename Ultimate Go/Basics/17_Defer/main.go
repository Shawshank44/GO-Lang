package main

import "fmt"

/*
	(i)In Go, the defer statement is used to delay the execution of a function until the surrounding function returns. It's commonly used for cleanup tasks such as closing files, unlocking mutexes, or logging.
	(ii)When a defer statement is encountered, Go schedules the function call to be run after the current function finishes—even if the function exits due to a panic.
	(iii)Deferred calls are executed in LIFO order (Last In, First Out).
*/

func main() {
	Process(1)
}

func Process(i int) {
	defer fmt.Println("Deffered i value : ", i) // runs last, prints value of i when defer is defined
	defer fmt.Println("First call")             // 4th to execute
	defer fmt.Println("Second call")            // 3rd to execute
	defer fmt.Println("Third call")             // 2nd to execute
	fmt.Println("Normal call")                  // 1st to execute
	i++
	fmt.Println("Actual i value : ", i) // executed immediately after increment
}

/*
	Explanation:
	📌 1. Why does Actual i value print after Normal call?
	Because it is written after the fmt.Println("Normal call") line.
	Execution happens in top-down order (like normal), so this part:
	executes in order:
	Print "Normal call"
	Increment i (now i = 2)
	Print "Actual i value : 2"

	 📌 2. Why does Deffered i value : show 1 and not 2?
	Because of how defer works:

	—Go evaluates the arguments immediately, but executes the function later.
	So at that moment, i was 1, and Go remembers:
	"Later I will run fmt.Println("Deffered i value : ", 1)"
	Even though i becomes 2 later, the deferred statement uses the old value (1), because it was captured early.

*/
