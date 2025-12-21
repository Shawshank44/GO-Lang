package main

import (
	"fmt"
	"reflect"
)

// type Person struct { // structs and fields
// 	Name string // fields must be always begin as capitalized
// 	Age  int
// }

type Greeter struct{}

func (G Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	// x := 42

	// v := reflect.ValueOf(x)

	// t := v.Type()

	// fmt.Println("Value :", v)
	// fmt.Println("Type : ", t)
	// fmt.Println("Kind : ", t.Kind())
	// fmt.Println("IsINT : ", t.Kind() == reflect.Int)
	// fmt.Println("IsSTRING : ", t.Kind() == reflect.String)
	// fmt.Println("Is Zero :", v.IsZero())

	// // Modifying the value using reflect package :
	// Y := 10
	// v1 := reflect.ValueOf(&Y).Elem()
	// v2 := reflect.ValueOf(&Y)
	// fmt.Println("Type Of v2 is :", v2.Type())
	// fmt.Println("Original value", v1.Int())
	// v1.SetInt(18)

	// fmt.Println("Modified value :", Y)

	// // Handling with Interface:
	// var itf interface{} = "Hello"
	// v3 := reflect.ValueOf(itf)
	// fmt.Println("v3 Type ", v3.Type())
	// if v3.Kind() == reflect.String {
	// 	fmt.Println("its a string value", v3.String())
	// }

	// Working with structs and fields:
	// p := Person{Name: "Alice", Age: 30}
	// v := reflect.ValueOf(p)
	// for i := range v.NumField() {
	// 	fmt.Printf("Field %d : %v \n", i, v.Field(i))
	// }

	// v1 := reflect.ValueOf(&p).Elem()

	// nameField := v1.FieldByName("Name")
	// if nameField.CanSet() {
	// 	nameField.SetString("Jane")
	// } else {
	// 	fmt.Println("Cannot change")
	// }

	// fmt.Println("Modified Person : ", p)

	// Working with Structs and Methods :
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method

	fmt.Println("Type:", t)
	for i := 0; i < t.NumMethod(); i++ {
		method = t.Method(i)
		fmt.Printf("Method: %d : %s \n", i, method.Name)
	}

	// now method holds the last method from the loop
	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("DOGE")})
	fmt.Println(results[0].String())
}
