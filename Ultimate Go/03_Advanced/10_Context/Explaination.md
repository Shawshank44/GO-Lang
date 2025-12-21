Here’s a structured guide to **Context** in Go:

---

## 1. Context — What It Is and Why It Exists

**Context** in Go is a way to carry deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines. It’s primarily used to control the lifecycle of concurrent operations and ensure they don’t leak or run indefinitely.

Purpose:

* Propagate cancellation signals across goroutines
* Enforce deadlines or timeouts
* Pass request-scoped values without global variables

Commonly used:

* HTTP request handling
* Database or API calls that need timeouts
* Long-running goroutines where cancellation may be required

---

## 2. Simple Code Example

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure resources are cleaned up

	go func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task completed")
		case <-ctx.Done():
			fmt.Println("Task canceled:", ctx.Err())
		}
	}(ctx)

	time.Sleep(4 * time.Second)
}
```

Key points:

* `context.WithTimeout` creates a context that automatically cancels after a duration
* `ctx.Done()` channel signals cancellation
* Always `defer cancel()` to release resources
* `ctx.Err()` provides the reason for cancellation (`context.DeadlineExceeded` or `context.Canceled`)

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Not calling `cancel()`

* Can cause resource leaks, especially with contexts that have timeouts

Fix:

* Always `defer cancel()` after creating a cancellable context

---

### Mistake 2: Ignoring `ctx.Done()`

* Long-running goroutines may continue running even if the parent is canceled

Fix:

* Regularly check `select { case <-ctx.Done(): ... }` in goroutines

---

### Mistake 3: Passing `nil` or background contexts everywhere

* Loses the ability to cancel operations or track deadlines

Fix:

* Use `context.Background()` only as a root context; pass derived contexts through function calls

---

## 4. Real-World Applications

### Scenario 1: HTTP Server Requests

* Pass context from HTTP handler to database queries so if the client disconnects, work is canceled

### Scenario 2: Microservice Calls

* Propagate cancellation or deadlines across services to avoid wasted computation or stuck processes

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a goroutine that prints numbers 1–5 with 1-second intervals. Use a context with timeout of 3 seconds to cancel it early if needed.

---

### Exercise 2 (Medium)

Write a function that performs a simulated API call with a 5-second delay. Use a context with deadline to cancel the function if it takes longer than 2 seconds.

---

### Exercise 3 (Hard)

Design a pipeline of two stages:

1. Stage 1 generates numbers 1–10
2. Stage 2 doubles the numbers
   Use a context to cancel the entire pipeline if processing takes more than 2 seconds, ensuring all goroutines exit cleanly.

---

## Thought-Provoking Question

In a large application with nested function calls and multiple goroutines, how would you design your use of contexts to **propagate deadlines and cancellation effectively** without creating tangled dependencies or over-canceling unrelated operations?
