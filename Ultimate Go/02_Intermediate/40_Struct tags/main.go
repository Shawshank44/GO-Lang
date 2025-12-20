package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"` // can use multple instances
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"` // omited completely even if value existss
}

func main() {
	person := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       36,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
