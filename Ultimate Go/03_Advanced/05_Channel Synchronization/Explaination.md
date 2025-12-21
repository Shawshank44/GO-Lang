Here’s a structured guide to **Channel Synchronization** in Go:

---

## 1. Channel Synchronization — What It Is and Why It Exists

**Channel synchronization** refers to using channels to coordinate the execution order of goroutines. Channels can act as **signals** to indicate when a goroutine has completed a task, allowing other goroutines to proceed safely.

Purpose:

* Coordinate multiple goroutines
* Avoid race conditions without explicit locks
* Wait for completion of tasks

Commonly used:

* Waiting for multiple goroutines to finish
* Sequential execution dependencies
* Implementing simple signaling or barriers

---

## 2. Simple Code Example

```go
package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true // Signal completion
}

func main() {
	done := make(chan bool)
	go worker(done)

	<-done // Wait for worker to finish
	fmt.Println("Worker finished")
}
```

Key points:

* Sending a value into a channel can signal completion
* Receiving from a channel can synchronize goroutines

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Forgetting to receive

* Sending on a channel without a corresponding receiver can block indefinitely

Fix:

* Always ensure a goroutine is ready to receive, or use buffered channels for non-blocking signaling

---

### Mistake 2: Closing channels incorrectly

* Closing a channel while goroutines may still send to it causes panic

Fix:

* Only the sender should close a channel, and only after all sends are done

---

### Mistake 3: Misusing channels for shared state

* Using a channel to store mutable shared data instead of just signaling can lead to complex code

Fix:

* Use channels primarily for signaling; use mutexes or other patterns for shared state

---

## 4. Real-World Applications

### Scenario 1: Waiting for multiple tasks

* Start several goroutines to process files; use a channel to signal completion before aggregating results

### Scenario 2: Sequential execution

* Ensure that certain initialization tasks finish before starting dependent goroutines

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a goroutine that prints a message and then sends a boolean to a channel. In the main goroutine, wait for the signal before printing "Done".

---

### Exercise 2 (Medium)

Start three goroutines that each sleep for a random duration and then send a message to a channel. The main goroutine should wait for all three messages before exiting.

---

### Exercise 3 (Hard)

Implement a fan-out/fan-in pattern:

* Launch 5 worker goroutines that perform some calculation and send results to an output channel
* Aggregate results in the main goroutine
* Ensure proper synchronization so no data is lost and all workers complete

---

## Thought-Provoking Question

Channels can synchronize goroutines without locks, but what are the **limitations of using channels for synchronization** in complex programs? Could relying solely on channels introduce performance bottlenecks or deadlocks, and how would you mitigate that?
