package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := `C:\Users\YourName\Documents\file.txt`

	// Join paths using filepath.join(as per your OS Go automatically creates the file paths)
	joinedPath := filepath.Join("downloads", "docs", "myVetris", "file.zip")
	fmt.Println("Joined path : ", joinedPath)

	// cleaning (correcting) the file path
	NormalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Cleaned path : ", NormalizedPath)

	// Splitting the file path
	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File: ", file)
	fmt.Println("Path: ", dir)
	fmt.Println(filepath.Base("/home/user/docs/file.txt")) // returns the last source of the file or directory

	// absolute or relative :
	fmt.Println("Relative : ", filepath.IsAbs(relativePath))
	fmt.Println("Absolute : ", filepath.IsAbs(absolutePath))

	// return the extension :
	fmt.Println(filepath.Ext(file))

	// Trimming the file path
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	// returning the relative path
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Returning the Absolute path :
	abspath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("absolute path : ", abspath)
}
