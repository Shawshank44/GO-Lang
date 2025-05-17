package main

import (
	"fmt"
)

type Person struct {
	firstName   string
	lastName    string
	age         int
	address     Address
	PhoneNumber // Anonymous struct embedding
}

// Struct embedding :
type Address struct {
	City    string
	Country string
}

// Anonymous struct embedding :
type PhoneNumber struct {
	OfficePH string
	CellPH   string
}

// Methods :
func (p Person) Fullname() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) IncrementAgeByOne() { // (*) Pointer helps to modify the existing value inside the struct
	p.age++
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "doe",
		age:       30,
		address: Address{
			City:    "London",
			Country: "England(Great Britain)",
		},
		PhoneNumber: PhoneNumber{
			OfficePH: "555-9876-000",
			CellPH:   "9876543210",
		},
	}

	p1 := Person{
		firstName: "Jane",
		lastName:  "Christine",
	}
	p1.address.City = "Newark(NJ)"
	p1.address.Country = "United states of America"

	p.IncrementAgeByOne()
	fmt.Println("Name :", p.Fullname(), "Age : ", p.age, "Address :", p.address.City, p.address.Country, "Phone Number :", p.OfficePH, p.CellPH)
	p1.IncrementAgeByOne()
	fmt.Println("Name :", p1.Fullname(), "Age : ", p1.age, "Address :", p1.address.City, p1.address.Country) // age will return 0 because no value assisgned

	// comparision of structs:
	fmt.Println(p == p1)

	//Annonymous structs
	// userad := struct {
	// 	userName string
	// 	email    string
	// }{
	// 	userName: "Jack1345",
	// 	email:    "Jack1345@google.com",
	// }

	// fmt.Println(userad.userName, userad.email)

}
