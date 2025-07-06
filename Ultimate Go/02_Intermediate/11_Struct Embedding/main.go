package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedded struct
	emp_ID string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hello My name is %s and Iam %d years old", p.Name, p.Age)
}

func (e Employee) introduce() { // overides person's introduce method
	fmt.Printf("Hello My name is %s and Iam %d years old and i earn %.2f", e.Name, e.Age, e.salary)
}

func main() {
	emp := Employee{
		Person: Person{Name: "John", Age: 30},
		emp_ID: "E001",
		salary: 5000.01,
	}

	fmt.Println("Name : ", emp.Name)
	fmt.Println("Name : ", emp.Age)
	fmt.Println("Name : ", emp.emp_ID)
	fmt.Println("Name : ", emp.salary)

	emp.introduce()
}
