Good. Methods on structs are where Go **quietly replaces “classes” without telling you**. If you misunderstand this, your designs will be bloated, inconsistent, and hard to reason about.

---

## 1. What Struct Methods Are (and why Go uses them)

A **method** is a function with a **receiver**.
When the receiver is a struct, the method is said to “belong” to that struct.

In Go:

* Methods define **behavior**
* Structs define **data**
* There is no inheritance—only composition

Purpose:

* Attach behavior to data
* Encapsulate logic related to a struct
* Enable interfaces

Struct methods are commonly used:

* To mutate or read struct state
* To define domain behavior
* To satisfy interfaces

Hard truth:

> Methods are about **ownership of behavior**, not object orientation.

---

## 2. Simple Code Example

```go
package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) Value() int {
	return c.value
}

func main() {
	c := Counter{}
	c.Increment()
	fmt.Println(c.Value())
}
```

Key points:

* Pointer receiver modifies state
* Value receiver reads state
* Go automatically handles `&` and `*` when calling methods

---

## 3. Common Beginner Mistakes (and how to avoid them)

### Mistake 1: Randomly mixing pointer and value receivers

Beginners pick receivers arbitrarily.

This causes:

* Confusing APIs
* Unexpected copying
* Interface issues

Fix:

* Use pointer receivers if the method mutates state
* Be consistent across the type

---

### Mistake 2: Putting logic in methods that don’t belong there

Some logic doesn’t “belong” to the struct.

Fix:

* Ask: “Is this behavior intrinsic to this data?”
* If not, use a standalone function or service

---

### Mistake 3: Assuming methods make code object-oriented

Go methods are not OOP magic.

Fix:

* Think in terms of data + behavior
* Use interfaces explicitly
* Avoid class-style thinking

---

## 4. Real-World Scenarios Where Struct Methods Matter

### Scenario 1: Domain Logic

Example:

* `Order.CalculateTotal()`
* `User.IsActive()`
* `Account.Withdraw()`

Behavior lives next to data.

---

### Scenario 2: Interface Implementation

Example:

* `Read()`, `Write()`
* `ServeHTTP()`

Methods allow structs to satisfy interfaces cleanly.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a struct that:

* Represents a rectangle
* Has methods to calculate area and perimeter

---

### Exercise 2 (Medium)

Create a struct that:

* Represents a bank account
* Has methods for deposit and withdrawal
* Prevents invalid operations

---

### Exercise 3 (Hard)

Design a small system where:

* Multiple structs have methods
* Interfaces are used to abstract behavior
* Pointer vs value receivers are chosen deliberately

This is about design, not syntax.

---

## Thought-Provoking Question (This separates good design from sloppy code)

Methods define behavior boundaries.

So answer this:
**How do you decide whether behavior should be a method on a struct, a standalone function, or part of a different type altogether—and how does that decision affect testability and coupling?**

If your answer is “whatever feels natural,” you’re not thinking deeply enough yet.
