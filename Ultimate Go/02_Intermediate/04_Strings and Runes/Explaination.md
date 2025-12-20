Good topic. Strings and runes expose **text reality vs. text illusion**. If you treat strings as “just characters,” you *will* write bugs—especially with non-English input. Let’s be precise.

---

## 1. What Strings and Runes Are (in Go, not fairy tales)

### Strings

In Go, a **string is a read-only slice of bytes**.

* Internally: UTF-8 encoded bytes
* Indexing a string gives you a **byte**, not a character
* Strings are immutable

Used when:

* Storing text
* Passing data
* Displaying or logging information

---

### Runes

A **rune** is an alias for `int32` and represents a **Unicode code point**.

* Used to represent a single logical character
* Required when working with non-ASCII text correctly

Used when:

* Iterating over characters (not bytes)
* Validating text
* Handling internationalization

Critical truth:

> Bytes ≠ characters ≠ runes

---

## 2. Simple Code Example

```go
package main

import "fmt"

func main() {
	s := "Go✓"

	fmt.Println(len(s)) // bytes

	for i, r := range s {
		fmt.Printf("index %d, rune %c\n", i, r)
	}
}
```

What this proves:

* `len(s)` counts bytes, not characters
* `range` decodes UTF-8 and yields runes
* Index `i` is a byte index, not character position

---

## 3. Common Beginner Mistakes (and how to avoid them)

### Mistake 1: Assuming `len(string)` means number of characters

It doesn’t.

Fix:

* Use `utf8.RuneCountInString`
* Or convert to `[]rune` if needed

---

### Mistake 2: Indexing strings like arrays of characters

`s[0]` gives a byte, not a rune.

Fix:

* Use `for range` to iterate runes
* Convert to `[]rune` when random access is required

---

### Mistake 3: Converting strings to `[]rune` unnecessarily

This allocates memory and slows things down.

Fix:

* Use `range` when possible
* Convert only when you truly need rune indexing

---

## 4. Real-World Scenarios Where This Matters

### Scenario 1: Internationalized Applications

* Usernames
* Messages
* Search and validation

Byte-based logic will break non-ASCII input.

---

### Scenario 2: Text Processing and Validation

* Character limits
* Sanitization
* Parsing user input

Counting bytes instead of characters causes subtle bugs.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Takes a string
* Prints its byte length and rune count

---

### Exercise 2 (Medium)

Write a function that:

* Iterates over a string
* Prints each rune and its Unicode code point

Test with non-ASCII characters.

---

### Exercise 3 (Hard)

Write a function that:

* Reverses a string correctly
* Works for Unicode characters
* Uses runes properly

This will expose whether you actually understand UTF-8.

---

## Thought-Provoking Question (This reveals real understanding)

UTF-8 makes strings efficient—but deceptive.

So answer this:
**When designing a Go API that processes text, how do you decide whether the API should operate on bytes, strings, or runes—and what bugs appear when you choose the wrong level of abstraction?**

If you just say “use runes everywhere,” you don’t understand the trade-off yet.
