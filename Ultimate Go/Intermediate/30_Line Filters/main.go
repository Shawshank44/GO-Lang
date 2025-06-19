package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Filter.txt")
	if err != nil {
		fmt.Println("Error in opening the file")
		return
	}

	defer func() {
		fmt.Println("Closing the file")
		file.Close()
	}()
	fmt.Println("file opened done")

	scanners := bufio.NewScanner(file)
	// Keyword to filter lines
	Keyword := "Important"
	// Read and filter lines
	linenumber := 1
	for scanners.Scan() {
		lines := scanners.Text()
		if strings.Contains(lines, Keyword) {
			Updatelines := strings.ReplaceAll(lines, Keyword, "Necessary")
			fmt.Printf("%d Filtered lines : %v \n", linenumber, lines)
			fmt.Printf("%d Updated lines : %v \n", linenumber, Updatelines)
			linenumber++
		}
	}

	err = scanners.Err()
	if err != nil {
		fmt.Println("error scanning file ", err)
	}

}
