package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// // Bufio reader
	// reader := bufio.NewReader(strings.NewReader("Hello, Bufio Packageee! \n"))
	// // Reading the data :
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes : %s", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Read string : ", line)

	// Bufio writer
	writer := bufio.NewWriter(os.Stdout)

	// writing the data:
	data := []byte("Hello, bufio package! \n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wrote %d bytes \n", n)

	// flust the buffer to ensure all data is return to stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	// writing with strings
	str := "This is a string. \n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes .\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
