package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct { // struct tags
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

type Address struct { // nested struct with json
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct { // storing the decoded JSON
	Emp_Name string `json:"employee_name"`
	Age      int    `json:"employee_age"`
	EmpId    string `json:"employee_id"`
	Role     string `json:"employee_role"`
}

func main() {
	person := Person{
		FirstName: "John",
		Age:       30,
		Email:     "John@mail.com",
		Address: Address{ // nesting with JSON
			City:  "Bengaluru",
			State: "Karnataka(KA)",
		},
	}
	// Using JSON (encoding)
	Jsonen, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(Jsonen))

	// Using JSON (decoding)
	JsonData := `{"employee_name":"Jen calvet","employee_age":35,"employee_id":"XYZ0065","employee_role":"AI engineer"}`

	var employee Employee
	err = json.Unmarshal([]byte(JsonData), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employee)

	// JSON and slices:
	ListOfCities := []Address{
		{City: "Newark", State: "NJ"},
		{City: "San Jose", State: "CA"},
		{City: "Atlanta", State: "GA"},
		{City: "Tampa", State: "FL"},
		{City: "Minneapolis", State: "MN"},
	}
	fmt.Println(ListOfCities)
	citylist, err := json.Marshal(ListOfCities)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(citylist))

	// Handling unknown structure data :
	UnJson := `{"name":"John Doe","age":30, "occupation":"VM Specs", "address":{"state":"Congo","country":"africa"}}`
	var data map[string]any
	err = json.Unmarshal([]byte(UnJson), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
