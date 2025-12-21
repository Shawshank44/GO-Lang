Here’s a structured guide to **Channels** in Go:

---

## 1. Channels — What They Are and Why They Exist

Channels are **conduits for communication between goroutines**. They allow one goroutine to send data to another safely, helping coordinate concurrent execution.

Purpose:

* Facilitate **synchronization** between goroutines
* Enable **safe data sharing** without explicit locks
* Support Go’s **communicating sequential processes (CSP)** model

Commonly used:

* Passing computed results from worker goroutines to a main goroutine
* Synchronizing tasks without using mutexes
* Building pipelines where data flows between stages concurrently

---

## 2. Simple Code Example

```go
package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine!"
}

func main() {
	messageChannel := make(chan string) // Create a channel

	go sendMessage(messageChannel)      // Launch goroutine

	message := <-messageChannel         // Receive from channel
	fmt.Println(message)
}
```

Key points:

* `make(chan Type)` creates a channel
* `<-` operator is used to **send** (`ch <- value`) and **receive** (`value := <- ch`) data
* Channels block until both sender and receiver are ready, providing **synchronization**

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Sending to a nil channel

* Sending or receiving on a nil channel **blocks forever**.

Fix:

* Always initialize channels with `make` before use

---

### Mistake 2: Forgetting to close a channel

* Not closing a channel when no more data is coming can lead to **deadlocks** during range loops

Fix:

* Use `close(channel)` when finished sending
* Only the sender should close a channel

---

### Mistake 3: Sending to a full buffered channel

* Sending to a buffered channel that is full blocks the sender indefinitely

Fix:

* Ensure the buffer has enough capacity, or use select statements for non-blocking sends

---

## 4. Real-World Applications

### Scenario 1: Worker pool

* Multiple worker goroutines process tasks, sending results back through a channel

### Scenario 2: Pipelines

* Data flows through multiple processing stages concurrently, each stage connected via channels

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a program where one goroutine sends 5 integers to the main goroutine via a channel, which prints each received value.

---

### Exercise 2 (Medium)

Implement two goroutines: one sends even numbers and the other sends odd numbers from 1 to 10 into a shared channel. The main goroutine should print all received numbers in the order they arrive.

---

### Exercise 3 (Hard)

Build a pipeline with three stages:

1. Generate numbers 1–10 in the first goroutine
2. Square each number in the second goroutine
3. Sum all squares in the third goroutine and send the final result to the main goroutine

---

## Thought-Provoking Question

Given that channels inherently synchronize goroutines, how would you decide when to use channels versus mutexes or other concurrency primitives in Go? Could overusing channels ever become a performance bottleneck, and why?
