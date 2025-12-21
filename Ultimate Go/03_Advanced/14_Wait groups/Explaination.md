Good. **WaitGroups are simple—and that’s exactly why people misuse them and create subtle bugs.** Let’s strip the fluff and get it right.

---

## 1. WaitGroups — What They Are and Why They Exist

A **WaitGroup** (`sync.WaitGroup`) is a synchronization primitive used to **wait for a collection of goroutines to finish**.

Purpose:

* Coordinate goroutines
* Block until all concurrent tasks are done
* Prevent premature program exit

Commonly used when:

* You launch multiple goroutines
* You don’t care about their return values (otherwise channels are better)
* You need a clean “wait until everything is done” point

**WaitGroups do ONE thing**: count goroutines.
They do **NOT** manage data, ordering, cancellation, or errors.

---

## 2. Simple WaitGroup Example

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "completed")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks finished")
}
```

Key facts:

* `Add()` increments the counter
* `Done()` decrements it
* `Wait()` blocks until counter hits zero
* `Done()` **must always be called**, even on errors

Miss one `Done()` and your program hangs forever.

---

## 3. Common Mistakes (And Why They’re Dangerous)

### Mistake 1: Calling `wg.Add()` inside the goroutine

This causes race conditions and undefined behavior.

Bad thinking:

> “It’s fine, it’s just one line”

Fix:

* Call `Add()` **before** starting the goroutine
* Treat `Add()` as setup, not execution

---

### Mistake 2: Forgetting `defer wg.Done()`

One early return → infinite wait → deadlock.

Fix:

* First line in the goroutine should be:

  ```go
  defer wg.Done()
  ```

If you don’t do this automatically, you’re writing fragile code.

---

### Mistake 3: Reusing a WaitGroup incorrectly

WaitGroups are not reset buttons.

Fix:

* A `WaitGroup` should represent **one lifecycle**
* Don’t reuse it across unrelated phases
* Create a new one if the meaning changes

---

## 4. Real-World Applications

### Scenario 1: Parallel Data Processing

* Read files concurrently
* Process each chunk independently
* Wait until all processing finishes before aggregating

---

### Scenario 2: Graceful Shutdown

* Multiple goroutines handling requests
* Signal shutdown
* Wait for all in-flight work to finish

Without this, you’re just killing work mid-flight.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Launch 5 goroutines that print numbers 1–5. Use a WaitGroup to ensure `main()` waits for all of them.

---

### Exercise 2 (Medium)

Start multiple goroutines that simulate work using `time.Sleep`. Ensure all complete before printing a final summary.

---

### Exercise 3 (Hard)

Build a program that:

* Starts N workers
* Uses a WaitGroup to wait for workers
* Combines WaitGroup with a channel to distribute jobs

If you try to replace the channel with a WaitGroup, you’re doing it wrong.

---

## Thought-Provoking Question

If WaitGroups only tell you **when work is done**, not **what happened**, **how would you design a system that handles errors, cancellation, and completion cleanly—without turning your code into spaghetti?**

That question separates “knows Go syntax” from “can design concurrent systems.”
