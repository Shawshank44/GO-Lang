package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// Extensible markup language

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Address Address
	// City    string   `xml:"city,omitempty"` // omiting the city
	// Email   string   `xml:"-"`              // omits even if the value exists
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

// assigning attributes
type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"` // will recognised as attribute of the parent element
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Pseudo  Pseudo
}

type Pseudo struct {
	Pseudo     xml.Name `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudo,attr"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			City:  "Oakland",
			State: "CA",
		},
	}

	// Encoding the XML
	XML, err := xml.MarshalIndent(person, "", "  ") // use MarshalIndent for better view
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(XML))

	// Decoding the XML
	xmlRaw := `
		<person>
			<name>Jane</name>
			<age>32</age>
			<Address>
				<city>ST Louis</city>
				<state>Mi</state>
			</Address>
		</person>
	`
	var personXML Person
	err = xml.Unmarshal([]byte(xmlRaw), &personXML)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(personXML)
	// fmt.Println(personXML.XMLName.Local)
	// fmt.Println(personXML.XMLName.Space)

	// assigning attributes
	book := Book{
		ISBN:   "567-573-689-636-878",
		Title:  "Go Bootcamp",
		Author: "Shashi",
		Pseudo: Pseudo{
			PseudoAttr: "PSP",
		},
	}

	bookxml, err := xml.MarshalIndent(book, "a", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bookxml))

}
