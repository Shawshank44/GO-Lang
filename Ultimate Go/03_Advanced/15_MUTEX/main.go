package main

import (
	"fmt"
	"sync"
)

// Basic Example :
// type Counter struct {
// 	mu    sync.Mutex // will get queued up
// 	count int
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock() // blocking in nature
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *Counter) GetValue() int {
// 	c.mu.Lock() // blocking in nature
// 	defer c.mu.Unlock()
// 	return c.count
// }

func main() {
	// // MUTEX AKA: Mutual exclusion
	// Basic Example :
	// var wg sync.WaitGroup
	// counter := &Counter{}
	// numGoroutines := 10

	// // wg.Add(numGoroutines) // both are same

	// for range numGoroutines {
	// 	wg.Add(1) // both are same
	// 	go func() {
	// 		defer wg.Done()
	// 		for range 1000 {
	// 			counter.Increment()
	// 		}
	// 	}()
	// }

	// wg.Wait()
	// fmt.Printf("Final counter value : %d \n", counter.GetValue())

	// Mutex without struct
	var counter int

	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	Increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go Increment()
	}

	wg.Wait()
	fmt.Printf("Final counter value : %d \n", counter)
}

/*
	In Go, a mutex (short for mutual exclusion) is a synchronization primitive used to protect shared resources from concurrent access. When multiple goroutines access or modify the same variable or data structure, a mutex ensures that only one goroutine at a time can access the critical section of code.

*/
