package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

/*
	sync.Pool, which is a concurrency-safe pool of temporary objects that can be reused to reduce allocation overhead and improve performance.
*/

func main() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Creating a new person.")
			return &Person{}
		},
	}

	// Get an object from the pool
	person := pool.Get().(*Person)
	person.name = "John"
	person.age = 18
	fmt.Println("Get Person : ", person)
	fmt.Printf("Name : %s, Age : %d \n", person.name, person.age)
	pool.Put(person)
	fmt.Println("Returned person to pool")

	person1 := pool.Get().(*Person)
	fmt.Println("Get Person again", person1)

	person2 := pool.Get().(*Person)
	fmt.Println("Get another person.", person2)

	// Returning object to the pool again
	pool.Put(person1)
	pool.Put(person2)
	fmt.Println("Returned Persons to pool")

	person3 := pool.Get().(*Person)
	fmt.Println("Get person : ", person3)
	person3.name = "Lupa"
	pool.Put(person3)

	person4 := pool.Get().(*Person)
	fmt.Println("Get Person : ", person4)
}
