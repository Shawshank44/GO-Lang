package main

import (
	"embed" // imported but not using directly (blank import)
	"fmt"
	"io/fs"
)

//go:embed example.txt
var content string

//go:embed baser
var baserfolder embed.FS

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Embeded content : ", content)
	content, err := baserfolder.ReadFile("baser/hello.txt")
	checkerror(err)

	fmt.Println("Embeded file content ", string(content))

	// listing the folder content :
	err = fs.WalkDir(baserfolder, "baser", func(path string, d fs.DirEntry, err error) error {
		checkerror(err)
		fmt.Println(path)
		return nil
	})
	checkerror(err)
}
