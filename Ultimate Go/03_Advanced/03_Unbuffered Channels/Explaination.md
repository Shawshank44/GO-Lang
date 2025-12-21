Here’s a structured guide to **Unbuffered Channels** in Go:

---

## 1. Unbuffered Channels — What They Are and Why They Exist

Unbuffered channels are channels **without any capacity**. Sending to an unbuffered channel **blocks** the sender until another goroutine receives the value, and receiving **blocks** until a value is sent.

Purpose:

* Provide **synchronous communication** between goroutines
* Ensure **data is handed off directly** from sender to receiver
* Implicitly synchronize execution without using explicit locks

Commonly used:

* Coordinating events between goroutines
* Passing a single value where immediate processing is required
* Enforcing strict ordering between producer and consumer

---

## 2. Simple Code Example

```go
package main

import "fmt"

func ping(ch chan string) {
	ch <- "Ping"  // Send blocks until main receives
}

func main() {
	channel := make(chan string) // Unbuffered channel

	go ping(channel)             // Launch goroutine

	message := <-channel         // Receive blocks until send happens
	fmt.Println(message)
}
```

Key points:

* `make(chan Type)` by default creates an **unbuffered channel**
* `<-` operator is used to send and receive
* Send and receive are **synchronous**, providing automatic coordination

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Ignoring blocking behavior

* Sending or receiving without a corresponding goroutine **blocks forever**, causing deadlock.

Fix:

* Always ensure a receiving goroutine exists when sending (and vice versa)

---

### Mistake 2: Using unbuffered channels when buffering is needed

* Sending multiple values in a loop can block if the receiver is slower

Fix:

* Use **buffered channels** if decoupling producer and consumer speed is necessary

---

### Mistake 3: Closing unbuffered channels prematurely

* Closing a channel while a send is still pending triggers a panic

Fix:

* Only close a channel when you’re certain all sends are complete

---

## 4. Real-World Applications

### Scenario 1: Task coordination

* One goroutine signals another to start processing, ensuring exact handoff

### Scenario 2: Synchronization points

* Use unbuffered channels to enforce execution order between critical goroutines

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program where a goroutine sends a greeting string to the main goroutine via an unbuffered channel. Main goroutine prints it.

---

### Exercise 2 (Medium)

Implement a “ping-pong” program: two goroutines send a message back and forth **5 times** using an unbuffered channel.

---

### Exercise 3 (Hard)

Create three goroutines representing stages of a pipeline. Use unbuffered channels to pass a number from stage 1 → stage 2 → stage 3. Each stage should double the number and pass it along. The main goroutine should print the final output.

---

## Thought-Provoking Question

Since unbuffered channels enforce synchronous handoff, how could overusing them in a highly concurrent program impact performance and throughput compared to buffered channels? In which scenarios would you **prefer unbuffered over buffered channels**, despite potential blocking?
