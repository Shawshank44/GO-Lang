package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	bufferSize = 5
)

type Buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *Buffer {
	b := &Buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (B *Buffer) Produce(Item int) {
	B.mu.Lock()
	defer B.mu.Unlock()
	for len(B.items) == bufferSize {
		B.cond.Wait()
	}
	B.items = append(B.items, Item)
	fmt.Println("Produced : ", Item)
	// B.cond.Signal() // used for single Go routines
	B.cond.Broadcast() // used for multiple Go routines
}

func (B *Buffer) Consume() int {
	B.mu.Lock()
	defer B.mu.Unlock()
	for len(B.items) == 0 {
		B.cond.Wait()
	}
	item := B.items[0]
	B.items = B.items[1:]
	fmt.Println("Consumed : ", item)
	// B.cond.Signal() // used for single Go routines
	B.cond.Broadcast() // used for multiple Go routines
	return item
}

func Producer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.Produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func Consumer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.Consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup
	wg.Add(2)
	go Producer(buffer, &wg)
	go Consumer(buffer, &wg)

	wg.Wait()
}
