package main

import (
	"fmt"
	"sync"
	"time"
)

type StatefulWorker struct {
	mu    sync.Mutex
	count int
	ch    chan int
}

func (SW *StatefulWorker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done() // Notify when the goroutine is done
		for value := range SW.ch {
			SW.mu.Lock()
			SW.count += value
			fmt.Println("current count:", SW.count)
			SW.mu.Unlock()
		}
	}()
}

func (SW *StatefulWorker) Send(value int) {
	SW.ch <- value
}

func main() {
	var wg sync.WaitGroup

	STW := &StatefulWorker{
		ch: make(chan int),
	}

	wg.Add(1)
	STW.Start(&wg)

	for i := 0; i < 5; i++ {
		STW.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

	close(STW.ch) // Close the channel to end the goroutine
	wg.Wait()     // Wait for the worker to finish
}
