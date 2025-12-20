package main

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Making a temporary file
	// Tempfile, err := os.CreateTemp("tempdir", "Tempfile")
	// CheckError(err)

	// fmt.Println(Tempfile.Name())
	// defer os.RemoveAll("tempdir")
	// defer Tempfile.Close()

	// Making a temporary directory :
	tempdir, err := os.MkdirTemp("", "GO course temp dir")
	CheckError(err)

	defer os.RemoveAll(tempdir)
	fmt.Println("Temporary directory created: ", tempdir)
}
