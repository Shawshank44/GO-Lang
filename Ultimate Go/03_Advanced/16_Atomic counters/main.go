package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (AC *AtomicCounter) Increment() {
	atomic.AddInt64(&AC.count, 1)
}

func (AC *AtomicCounter) Getvalue() int64 {
	return atomic.LoadInt64(&AC.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter value : %d", counter.Getvalue())
}
