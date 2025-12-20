Good. Random numbers are another area where beginners confidently do the **wrong thing** and only realize it much later. Let’s get this straight.

---

## 1. What Random Numbers Are (Purpose & When to Use Them)

In Go, “random numbers” usually come from two places:

* `math/rand` → **pseudo-random**, fast, deterministic
* `crypto/rand` → **cryptographically secure**, slower, non-deterministic

Purpose:

* Simulate randomness
* Shuffle data
* Make decisions unpredictable
* Generate IDs, tokens, or test data (depending on source)

Commonly used when:

* Games and simulations
* Load testing
* Shuffling or sampling
* Security-sensitive operations (only with `crypto/rand`)

Hard rule:

> If security matters and you use `math/rand`, your code is wrong.

---

## 2. Simple Code Example

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the generator
	rand.Seed(time.Now().UnixNano())

	// Random number between 0 and 9
	n := rand.Intn(10)
	fmt.Println(n)
}
```

What matters:

* `math/rand` is **deterministic**
* Without seeding, you get the same sequence every run
* `Intn(n)` returns `[0, n)`

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Forgetting to seed the RNG

Classic rookie mistake.

Symptom:

* Same “random” numbers every run

Fix:

* Seed once at program startup
* Use `time.Now().UnixNano()`

---

### Mistake 2: Thinking pseudo-random means “truly random”

It doesn’t.

Reality:

* `math/rand` produces predictable sequences
* Anyone with the seed can reproduce results

Fix:

* Use `crypto/rand` for passwords, tokens, keys
* Use `math/rand` only for non-security logic

---

### Mistake 3: Reseeding repeatedly

Some beginners reseed inside loops.

Why it’s bad:

* Reduces randomness
* Can produce repeated values

Fix:

* Seed **once**
* Never reseed unless you fully understand why

---

## 4. Real-World Applications Where Random Numbers Matter

### Scenario 1: Simulations and Games

* Dice rolls
* Enemy behavior
* Random events

Predictability is acceptable—and sometimes desirable.

---

### Scenario 2: IDs, Tokens, and Security

* Session tokens
* Password reset links
* API keys

Here, weak randomness is a **security vulnerability**, not a bug.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Generates a random number within a given range
* Prints it multiple times

---

### Exercise 2 (Medium)

Create a function that:

* Simulates rolling two dice
* Returns the result
* Uses randomness correctly

---

### Exercise 3 (Hard)

Build a small utility that:

* Generates random strings
* Allows configurable length
* Clearly distinguishes between secure and non-secure randomness

Design matters more than output.

---

## Thought-Provoking Question

Most developers use random numbers without understanding their guarantees.

**How can using the wrong type of randomness silently turn a “working” system into a security or correctness disaster, even though all tests pass?**

If your answer is “that won’t happen to me,” you’re exactly the target audience for that disaster.
