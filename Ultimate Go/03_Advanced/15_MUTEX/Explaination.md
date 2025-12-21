Good. **Mutex is where people stop “playing with goroutines” and start breaking programs.** If you don’t understand *why* you’re locking, you shouldn’t be locking at all.

---

## 1. Mutex — What It Is and Why It Exists

A **Mutex** (`sync.Mutex`) enforces **mutual exclusion**:
only **one goroutine** can access a critical section at a time.

Purpose:

* Protect shared mutable state
* Prevent race conditions
* Ensure data consistency

Used when:

* Multiple goroutines **read/write the same data**
* You cannot redesign the data flow to avoid sharing
* Performance matters more than architectural purity

Blunt truth:

> If you don’t have shared mutable data, you don’t need a mutex.
> If you do have it and don’t lock, your program is wrong.

---

## 2. Simple Mutex Example

```go
package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
```

What matters:

* `Lock()` before touching shared state
* `Unlock()` immediately after
* Critical section must be **as small as possible**

---

## 3. Common Mistakes (These Cause Deadlocks or Corruption)

### Mistake 1: Forgetting to unlock

This **will deadlock your program**.

Fix:

```go
mu.Lock()
defer mu.Unlock()
```

If you don’t use `defer` by default, you’re writing unsafe code.

---

### Mistake 2: Locking too much code

People lock entire functions “just to be safe.” That kills concurrency.

Fix:

* Lock only the lines that touch shared data
* Do NOT perform I/O, sleep, or network calls while holding a lock

---

### Mistake 3: Using Mutex where Channels fit better

Mutexes hide data flow and increase cognitive load.

Fix:

* If goroutines communicate by passing data → use channels
* If goroutines share memory → use mutex
* If you can avoid sharing → do that instead

---

## 4. Real-World Applications

### Scenario 1: In-Memory Cache

* Multiple goroutines reading/writing a map
* Mutex ensures map safety and consistency

Without a mutex, your program will panic or corrupt memory.

---

### Scenario 2: Shared Metrics / Counters

* Request counters
* Error rates
* Concurrent statistics updates

This is a classic mutex use-case: small, fast, critical sections.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a shared integer. Start multiple goroutines that increment it safely using a mutex.

---

### Exercise 2 (Medium)

Protect a shared `map[string]int` where goroutines update values concurrently.

---

### Exercise 3 (Hard)

Build a thread-safe struct that:

* Encapsulates its own mutex
* Exposes methods that safely read and write internal state
* Never exposes the raw data directly

If you expose the map, you failed the exercise.

---

## Thought-Provoking Question

**If a mutex guarantees correctness but reduces parallelism, how do you decide whether correctness or throughput matters more—and what would you change in your design to avoid choosing at all?**

That question forces you to think like a system designer, not just a Go programmer.
