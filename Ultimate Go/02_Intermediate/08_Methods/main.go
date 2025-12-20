package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}
type Shape struct {
	Rectangle
}

// Value receiver method
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Pointer receiver method
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int // custom datatype

func (m MyInt) IsPositive() bool {
	return m > 0
}

func main() {
	rectangle := Shape{
		Rectangle: Rectangle{
			length: 10,
			width:  9,
		},
	}

	fmt.Println("Area of the rectangle is :", rectangle.Area())
	rectangle.Scale(2)
	fmt.Println("Area of the rectangle is :", rectangle.Area())

	// Customer datatype
	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
}
