Here’s a structured guide to **Go Routines** in Go:

---

## 1. Go Routines — What They Are and Why They Exist

Go routines are **lightweight threads managed by the Go runtime**, allowing concurrent execution of functions.

Purpose:

* Enable **concurrent programming** without manually managing threads
* Efficiently utilize CPU resources
* Simplify asynchronous and parallel tasks

Commonly used:

* Performing tasks simultaneously (e.g., I/O operations, API calls, data processing)
* Building concurrent servers or applications
* Avoiding blocking in programs that handle multiple operations

---

## 2. Simple Code Example

```go
package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	go greet("Alice") // Runs concurrently
	go greet("Bob")

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function ends")
}
```

Key points:

* Use the `go` keyword before a function call to launch a goroutine
* Goroutines run **asynchronously** and may finish at different times
* Use synchronization (channels, WaitGroups) to manage completion

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Exiting main before goroutines complete

* Main function may end before goroutines finish, terminating the program

Fix:

* Use `time.Sleep` (temporary), or better, `sync.WaitGroup` to wait for completion

---

### Mistake 2: Concurrent access to shared variables

* Multiple goroutines modifying the same variable can cause **race conditions**

Fix:

* Use channels or `sync.Mutex` to safely share data

---

### Mistake 3: Forgetting that goroutines are lightweight but not free

* Launching thousands without control can lead to resource exhaustion

Fix:

* Limit concurrent goroutines using buffered channels, worker pools, or semaphores

---

## 4. Real-World Applications

### Scenario 1: Web servers

* Handle multiple client requests concurrently without blocking other connections

### Scenario 2: Parallel data processing

* Downloading files or processing large datasets simultaneously to improve performance

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a program that launches three goroutines, each printing a different message, and ensure all messages are printed before the program exits.

---

### Exercise 2 (Medium)

Write a program that launches multiple goroutines to calculate the square of numbers from 1 to 5 and sends results through a channel to the main function, which prints them.

---

### Exercise 3 (Hard)

Implement a worker pool with 5 workers using goroutines to process a list of 20 tasks concurrently, ensuring all tasks are completed before the program exits.

---

## Thought-Provoking Question

If goroutines are so lightweight and easy to create, why can launching too many of them still lead to performance problems or crashes? How would you design a Go program to balance concurrency and resource usage effectively?
