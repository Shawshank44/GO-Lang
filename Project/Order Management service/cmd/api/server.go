package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Order management")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", home)

	fmt.Printf("Server started successfully on http://localhost%s", os.Getenv("API_PORT"))

	err = http.ListenAndServe(os.Getenv("API_PORT"), nil)
	if err != nil {
		log.Fatalln("Server error : ", err)
	}
}
