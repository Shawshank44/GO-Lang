package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Accessing the environment variables
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var : ", user)
	fmt.Println("Home env var : ", home)

	// Setting the Env variable
	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("fruit env variable is : ", os.Getenv("FRUIT"))

	// Unsetting the Env varable :
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("fruit env variable is : ", os.Getenv("FRUIT"))

	// for _, e := range os.Environ() { // os.Environ returns all the Environment variables present in your system
	// 	keypair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(keypair[0])
	// }

	// strings.SplitN example
	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
}
