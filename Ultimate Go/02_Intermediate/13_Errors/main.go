package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error: square root of negative number is not possible")
	}

	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}
	return nil
}

// custom errors
type MyError struct {
	message string
}

func (m *MyError) Error() string { // error interface of built in package
	return fmt.Sprintf("Error : %s", m.message)
}

func errprocessor() error {
	return &MyError{"custom Error message"}
}

func readConfig() error {
	return errors.New("Config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData : %w", err)
	}
	return nil
}

func main() {

	// res, err := sqrt(-13)

	// if err != nil { // usually used
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)

	// data := []byte{}

	// if err := process(data); err != nil { // shorthand notation
	// 	fmt.Println("Error : ", err)
	// 	return
	// }
	// fmt.Println("Data processed Successfully")

	// errs := errprocessor()

	// if errs != nil {
	// 	fmt.Println(errs)
	// 	return
	// }

	errsf := readData()

	if errsf != nil {
		fmt.Println(errsf)
		return
	}
	fmt.Println("data read successfully")
}
