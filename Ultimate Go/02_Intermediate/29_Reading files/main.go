package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("Readstring.txt")
	// if err != nil {
	// 	fmt.Println("Error in creating a file : ", err)
	// }
	// defer file.Close()

	// _, err = file.WriteString("Hello Go!\n")
	// if err != nil {
	// 	fmt.Println("Error writing to file")
	// 	return
	// } else {
	// 	fmt.Println("write string file updated")
	// }

	file, err := os.Open("Readstring.txt")
	if err != nil {
		fmt.Println("error Opening file :", err)
		return
	}
	defer func() {
		fmt.Println("Closing file")
		file.Close()
	}()
	fmt.Println("file Opened successfully")

	// Reading the contents of the opened file :
	// data := make([]byte, 1024) // buffer to read data
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("error reading the data file : ", err)
	// 	return
	// }
	// fmt.Println("file content : ", string(data))

	scanners := bufio.NewScanner(file)
	for scanners.Scan() {
		line := scanners.Text()
		fmt.Println("Line : ", line)
	}
	err = scanners.Err()
	if err != nil {
		fmt.Println("error reading file : ", err)
		return
	}
}
