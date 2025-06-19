package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("error in creating a file")
		return
	}
	defer file.Close()

	// writing some data to a file (Using bytes) :
	data := []byte("Hello World\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("error while writing to the file")
		return
	} else {
		fmt.Println("Data has been witten in the file")
	}

	// writing some data to a file (Using write string)
	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error in creating a file : ", err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go!\n")
	if err != nil {
		fmt.Println("Error writing to file")
		return
	} else {
		fmt.Println("write string file updated")
	}

}
