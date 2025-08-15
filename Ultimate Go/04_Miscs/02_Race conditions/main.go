package main

import (
	"fmt"
	"sync"
)

// go run -race main.go - use this command to know race conditions

func main() {
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
