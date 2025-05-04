package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Defered function")
	fmt.Println("Starting the main function")

	os.Exit(1)
	fmt.Println("End of main function")
}

/*
	1. The `main()` function starts and registers a deferred call to `fmt.Println("Defered function")`.
	2. It prints `"Starting the main function"`.
	3. `os.Exit(1)` is called, which immediately terminates the program with exit code 1.
	4. Because `os.Exit` exits the program instantly, the deferred function is **not** executed.
	5. The line `fmt.Println("End of main function")` is also never reached.
*/
