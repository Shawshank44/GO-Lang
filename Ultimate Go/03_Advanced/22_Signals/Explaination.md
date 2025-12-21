Good topic. Signals separate people who write **programs** from people who run **systems**. Here’s the straight version, no fluff.

---

## 1. What Signals Actually Are (and Why They Exist)

**Signals** are OS-level notifications sent to a running process to tell it that **something happened**.

In Go, signals are mainly used to:

* Gracefully shut down programs
* Handle interrupts (Ctrl+C)
* React to kill, terminate, or reload events

Commonly used when:

* Running long-lived services (servers, workers, daemons)
* Managing resources (files, DB connections, goroutines)
* Deployments, restarts, or container orchestration

Hard truth:

> If your program ignores signals, it’s not production-ready.

---

## 2. Simple Code Example (Signal Handling)

This example listens for an interrupt and shuts down cleanly.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Program running. Press Ctrl+C to stop.")

	sig := <-sigChan
	fmt.Println("Received signal:", sig)

	fmt.Println("Cleaning up before exit...")
}
```

What matters here:

* `signal.Notify` registers interest in specific signals
* The channel **must be buffered**
* Blocking on `<-sigChan` waits until a signal arrives

---

## 3. Common Mistakes (and Why They Break Systems)

### Mistake 1: Using unbuffered signal channels

Beginners often write `make(chan os.Signal)`.

Why this is wrong:

* Signals can be dropped
* Delivery is asynchronous
* Runtime may block trying to send

Avoidance:

* Always use a buffered channel (`size >= 1`)

---

### Mistake 2: Assuming all signals are catchable

People try to handle `SIGKILL` or `SIGSTOP`.

Reality:

* Some signals **cannot** be caught or ignored
* The OS will terminate your process anyway

Avoidance:

* Learn which signals are catchable (`SIGINT`, `SIGTERM`)
* Design for forced termination

---

### Mistake 3: Handling signals but not stopping goroutines

The program “exits” but goroutines are still running.

Why this happens:

* Signal handler prints a message and exits
* No coordination with workers

Avoidance:

* Use context cancellation
* Signal → cancel context → wait → exit

---

## 4. Real-World Applications

### Scenario 1: Graceful Server Shutdown

* Stop accepting new requests
* Finish in-flight requests
* Close DB connections
* Exit cleanly

Signals are the **only sane trigger** for this.

---

### Scenario 2: Container & Kubernetes Lifecycle

* Kubernetes sends `SIGTERM`
* You have limited time to shut down
* Ignoring signals = forced kill

If you write Go services and ignore signals, Kubernetes will punish you.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Listens for `SIGINT`
* Prints a message
* Exits cleanly

Focus on correct signal registration.

---

### Exercise 2 (Medium)

Create a program that:

* Runs a background goroutine printing a message every second
* Stops the goroutine when a termination signal is received
* Ensures clean exit without leaks

This tests coordination, not syntax.

---

### Exercise 3 (Hard)

Build a service that:

* Handles `SIGTERM`
* Gracefully shuts down multiple workers
* Times out if shutdown takes too long
* Logs shutdown progress

This is real production behavior.

---

## Thought-Provoking Question

**If your program receives a termination signal while it’s halfway through a critical operation, how do you decide whether to finish the work, roll it back, or abort immediately—and where should that decision live in your code?**

If you can’t answer that clearly, you don’t control your shutdown—you’re just hoping.
