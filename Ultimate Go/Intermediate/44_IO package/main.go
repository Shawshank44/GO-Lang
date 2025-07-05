package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}

func closeResourse(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // this creates a memory on the stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World")
	MR := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // this allocates memory on the heap
	_, err := buf.ReadFrom(MR)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func WriteToFile(fpath string, data string) {
	file, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer closeResourse(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}

	// other method to write a code
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

type MyClose struct { // interface example
	name string
}

func (m MyClose) Close() error {
	fmt.Println("Closing : ", m.name)
	return nil
}

func main() {
	// Reader from Reader :
	readFromReader(strings.NewReader("Hello Readers!"))
	// Write from Writer :
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writer")
	fmt.Println(writer.String())
	// Buffer Example :
	bufferExample()
	// Multi Reader Example :
	multiReaderExample()
	//Pipe Example:
	pipeExample()
	WriteToFile("Io.txt", "Hello IO package")

	// close Interface example :
	res := &MyClose{
		name: "Tester",
	}
	closeResourse(res)

}
