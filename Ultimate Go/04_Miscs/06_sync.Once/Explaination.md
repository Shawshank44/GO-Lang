Good. `sync.Once` looks trivial on the surface, which is exactly why people misuse it. Let’s strip it down and be precise.

---

## 1. `sync.Once` — What It Is and When to Use It

`sync.Once` guarantees that **a piece of code executes exactly once**, even when multiple goroutines try to run it concurrently.

Key properties:

* Thread-safe
* Zero explicit locking required by the caller
* Once it runs successfully, it will **never run again**

When it’s commonly used:

* One-time initialization (config, logger, DB connection)
* Lazy initialization shared across goroutines
* Avoiding double-init race conditions

When it’s **not** appropriate:

* When you need to re-run logic
* When failure should allow retry
* When behavior depends on runtime conditions

Hard truth:

> `sync.Once` is for **initialization**, not control flow.

---

## 2. Simple Code Example Demonstrating `sync.Once`

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initResource() {
	fmt.Println("Initializing resource")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initResource)
	fmt.Println("Worker running")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
}
```

What this proves:

* `initResource()` runs exactly once
* All goroutines safely share the result
* No race, no manual locks

---

## 3. Common Mistakes & How to Avoid Them

### Mistake 1: Assuming `sync.Once` retries after failure

This is a **critical misunderstanding**.

If the function panics:

* `Once` considers it **done**
* It will never run again

Avoidance:

* Make the function panic-safe
* Or handle retry logic outside `sync.Once`

Rule:

> `sync.Once` does not mean “until it works”.

---

### Mistake 2: Using `sync.Once` for conditional logic

People abuse it like this:

```go
if condition {
	once.Do(fn)
}
```

Problem:

* Condition may change later
* `Once` permanently locks behavior

Avoidance:

* Evaluate conditions **inside** the `Do` function
* Or don’t use `Once` at all

---

### Mistake 3: Creating multiple `sync.Once` instances accidentally

Example:

```go
func get() {
	var once sync.Once
	once.Do(init)
}
```

This does nothing useful.
Each call creates a **new Once**.

Avoidance:

* Store `sync.Once` at package or struct level
* Scope matters

---

## 4. Real-World Applications Where `sync.Once` Is Useful

### Scenario 1: Lazy configuration or environment loading

* Config loaded only when first needed
* Safe across many goroutines
* No startup penalty

---

### Scenario 2: Singleton-style resource initialization

Examples:

* Logger setup
* Metrics exporter
* Connection pool bootstrap

Note:

> Go discourages singletons—but sometimes they’re unavoidable.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program where:

* Multiple goroutines attempt to initialize a global variable
* Initialization logic must run exactly once

Focus: understanding shared `sync.Once` scope.

---

### Exercise 2 (Medium)

Build a struct that:

* Lazily initializes an internal resource
* Multiple goroutines call a method concurrently
* Resource must be initialized once and shared

Focus: embedding `sync.Once` inside a struct.

---

### Exercise 3 (Hard)

Design a system where:

* Initialization may fail
* You must detect failure
* You must decide whether `sync.Once` is still appropriate

Focus: recognizing when **not** to use `sync.Once`.

---

## Thought-Provoking Question

**If `sync.Once` permanently disables retries—even after a panic—what does that tell you about the kind of problems it was designed to solve, and which problems it explicitly refuses to help you with?**

Think carefully. That design choice is intentional, not a limitation.
