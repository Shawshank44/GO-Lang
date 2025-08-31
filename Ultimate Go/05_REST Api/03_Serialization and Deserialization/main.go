package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Marshall and unmarshall
	user := User{Name: "Alice", Email: "Alice@mail.com"}
	fmt.Println(user)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unmarshalled data : ", user1)

	// Decode and Encode
	Data := `{"name" : "Josh", "email": "josh@email.com"}`

	reader := strings.NewReader(Data)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user2)

	// encode :
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded Json String", buf.String())

	// Handling unknown Json data with Decoder and encoder
	JSOND := `{"name" : "Jack", "age" : "30", "email" : "jack@gmail.com"}`
	var MAPdata map[string]any
	Ureader := strings.NewReader(JSOND)
	Udecode := json.NewDecoder(Ureader)
	err = Udecode.Decode(&MAPdata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(MAPdata["name"])

}
