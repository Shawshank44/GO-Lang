package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine 1 finished")
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine 2 finished")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function complete")
	// select {}
}

/*
	A deadlock in programming is a situation where two or more processes (or threads) are stuck waiting for resources in such a way that none of them can proceed — kind of like two people holding doors open for each other, both insisting the other should go first, and neither actually moving.

	How Deadlocks Happen

	In multithreaded or multiprocessing systems, a deadlock usually occurs when all of these conditions are met (Coffman’s conditions):

	Mutual exclusion – A resource can only be used by one process/thread at a time.

	Hold and wait – A thread is holding one resource and waiting to acquire another.

	No preemption – A resource can’t be forcibly taken away; it must be released voluntarily.

	Circular wait – There’s a circular chain where each process waits for a resource held by the next.

	How to Prevent Deadlocks

	You can avoid deadlocks by breaking at least one of the Coffman conditions:

	Lock ordering – Always acquire resources in a predetermined global order.

	Lock timeout – Use try_lock or timeouts so a thread can give up instead of waiting forever.

	Deadlock detection – Regularly check for circular waits and recover (e.g., abort a process).

	Avoid unnecessary locks – Minimize shared mutable state.

*/
