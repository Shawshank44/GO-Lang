package main

import (
	"fmt"
	"sync"
)

var RWMU sync.RWMutex
var COUNTER int

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	RWMU.RLock()
	fmt.Println("Read counter : ", COUNTER)
	RWMU.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	RWMU.Lock()
	COUNTER = value
	fmt.Printf("Written value %d for the counter", value)
	RWMU.Unlock()
}

func main() {
	// RW Mutex - Read write Mutex
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	// time.Sleep(time.Second)
	go writeCounter(&wg, 18)

	wg.Wait()
}

/*
	An RWMutex (short for Read-Write Mutex) is a synchronization primitive that allows concurrent read access but exclusive write access to a shared resource.

	It’s most commonly seen in multithreaded programming (for example, in Go’s sync.RWMutex or similar constructs in other languages).

	How it works

	Multiple readers can hold the lock at the same time — as long as no writer has it.

	Only one writer can hold the lock — and when writing, it blocks all readers and other writers.

	Fairness rules vary by implementation — some RWMutex implementations may prefer writers, others readers, to prevent starvation.

	Typical operations

	Read lock

	Acquired with something like RLock() (Go) or lock_shared() (C++).

	Many readers can hold this at once.

	Blocks if a writer currently holds the lock.

	Read unlock

	Releases the read lock (RUnlock()).

	Write lock

	Acquired with something like Lock() (Go) or lock() (C++).

	Only one writer can hold it at a time, and no readers are allowed during this period.

	Write unlock

	Releases the write lock (Unlock()).

	Why use it?

	Performance: If reads vastly outnumber writes, RWMutex allows higher concurrency than a regular Mutex, since readers don’t block each other.

	Safety: Ensures consistent access to shared state, preventing data races.


*/
