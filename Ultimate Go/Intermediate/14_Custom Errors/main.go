package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	code    int
	message string
	err     error
}

func (e *CustomError) Error() string { // inbuilt Error() Interface
	return fmt.Sprintf("Error %d : %s", e.code, e.message)

}

// function that returns a custom error
// func doSomething() error {
// 	return &CustomError{
// 		code:    500,
// 		message: "Something went wrong",
// 	}
// }

func doSomethingElse() error {
	return errors.New("Internal error")
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &CustomError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}
	return nil
}

func main() {
	err := doSomething()

	if err != nil {
		fmt.Println(err)
		return // we return here becuase if occurs we don't want to execute rest of the statement
	}
	fmt.Println("Operation succesfull")

}
