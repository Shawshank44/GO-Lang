Good. Formatting verbs look trivial, but misuse them and your output becomes **misleading, hard to debug, or outright wrong**. This is about precision, not cosmetics.

---

## 1. What Formatting Verbs Are (and why they matter)

Formatting verbs are **placeholders used by `fmt` functions** (`Printf`, `Sprintf`, `Fprintf`) to control how values are displayed.

Examples:

* `%d` → integers
* `%f` → floating-point numbers
* `%s` → strings
* `%v` → default format
* `%T` → type of the value

Purpose:

* Format output precisely
* Control readability and correctness
* Avoid ambiguous or misleading output

Formatting verbs are commonly used:

* Logging
* Debugging
* CLI output
* Reports and diagnostics

Hard truth:

> Bad formatting hides bugs and confuses users.

---

## 2. Simple Code Example

```go
package main

import "fmt"

func main() {
	name := "Go"
	version := 1.22
	lines := 150

	fmt.Printf("Language: %s\n", name)
	fmt.Printf("Version: %.2f\n", version)
	fmt.Printf("Lines of code: %d\n", lines)
	fmt.Printf("Type of version: %T\n", version)
}
```

Key takeaways:

* `%s`, `%f`, `%d` are type-specific
* Precision (`.2`) matters
* `%T` is invaluable for debugging

---

## 3. Common Beginner Mistakes (and how to avoid them)

### Mistake 1: Using `%v` everywhere

Yes, it works. No, it’s not good enough.

Fix:

* Use explicit verbs for clarity
* Reserve `%v` for quick debugging

---

### Mistake 2: Mismatching verbs and types

This produces wrong output or runtime errors.

Example:

```go
fmt.Printf("%d", "text")
```

Fix:

* Know the type you’re formatting
* Use `%T` when unsure

---

### Mistake 3: Ignoring width and precision

Default formatting is often ugly or misleading.

Fix:

* Use width (`%6d`) and precision (`%.2f`)
* Align output intentionally

---

## 4. Real-World Scenarios Where Formatting Verbs Matter

### Scenario 1: Logging and Debugging

Logs must be:

* Clear
* Consistent
* Machine- and human-readable

Formatting mistakes waste debugging time.

---

### Scenario 2: CLI Tools and Reports

* Tables
* Summaries
* Status output

Poor formatting makes tools painful to use.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Prints an integer, float, and string
* Uses correct formatting verbs for each

---

### Exercise 2 (Medium)

Write a program that:

* Prints a table of numbers
* Aligns columns using width formatting

---

### Exercise 3 (Hard)

Write a program that:

* Logs structured data
* Prints values, types, and memory addresses
* Uses multiple formatting verbs correctly

This forces deliberate formatting choices.

---

## Thought-Provoking Question (This goes beyond syntax)

Formatting controls perception.

So answer this:
**How can careless use of formatting verbs hide logical errors in a program, and what formatting choices help expose bugs instead of masking them?**

If you think formatting is “just output,” you’re missing its diagnostic power.
