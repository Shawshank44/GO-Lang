package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64 // it is compulsary to have both methods if implemented
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// func MyPrinter(i ...interface{}) {
// 	for _, v := range i {
// 		fmt.Println(v)
// 	}
// }

func MyPrinter(i ...any) { // any keyword is shorthand for interface
	for _, v := range i {
		fmt.Println(v)
	}
}

func TypeCast(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type is Integer")
	case string:
		fmt.Println("Type is String")
	case float64:
		fmt.Println("Type is Float")
	case bool:
		fmt.Println("Type is Boolean")
	}
}

func main() {
	// r := Rect{width: 3, height: 4}
	// c := Circle{radius: 5}

	// measure(r)
	// measure(c)

	// MyPrinter(1, "Strings", 45.9, true)
	TypeCast(true)
}
