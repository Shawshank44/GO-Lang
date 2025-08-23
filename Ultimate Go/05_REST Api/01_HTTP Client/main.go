package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create a new http client
	client := &http.Client{}
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/2")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}
