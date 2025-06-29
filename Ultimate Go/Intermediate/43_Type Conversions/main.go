package main

import (
	"fmt"
)

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool(c) // Cannot be converted
	e := 3.14
	f := int(e)
	fmt.Printf("The type of a is %T \n", a)
	fmt.Printf("The type of b is %T \n", b)
	fmt.Printf("The type of c is %T \n", c)
	fmt.Printf("The type of e is %T \n", e)
	fmt.Printf("The type of f is %T \n", f)

	// Strings and bytes
	g := "Hello"
	var h []byte
	h = []byte(g)
	fmt.Printf("The g is %T \n", h)

	byts := []byte{72, 120} // can only take upto 255
	fmt.Println(string(byts))
}
