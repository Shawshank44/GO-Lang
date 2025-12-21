package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func Initialize() {
	fmt.Println("This will not be repeated no matter how many time we call this function using once.do")
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Go Routine %d \n", i)
			once.Do(Initialize) // this will only run once no matter how many Go routines are there to execute.
		}()
	}

	wg.Wait()
}
