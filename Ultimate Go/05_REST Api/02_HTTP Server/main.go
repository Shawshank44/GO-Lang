package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello Server")
	})

	const PORT string = "127.0.0.1:3000"
	// const PORT string = ":3000" // we can enter like this as well

	fmt.Println("Server Listening on Port : ", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
