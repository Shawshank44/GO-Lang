Good topic. Conditional statements are where **logic actually branches**. If you misunderstand this, your program may compile perfectly and still behave like an idiot.

I’ll keep this direct.

---

## 1. What Conditional Statements Are (and why they exist)

In Go, **conditional statements let your program choose different execution paths based on conditions**.

Their purpose is simple:

* **Make decisions**
* **Enforce rules**
* **Validate state**
* **Control flow**

You use conditionals whenever **the program should not behave the same way all the time**. If everything runs linearly, you’re not writing software—you’re writing a script.

Go mainly uses:

* `if`, `else if`, `else`
* `switch`

---

## 2. Simple Go Example

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("Eligible")
    } else {
        fmt.Println("Not eligible")
    }
}
```

Key points:

* Conditions **must be boolean**
* No parentheses around conditions (this isn’t C or Java)
* Blocks are mandatory (`{}`)

---

## 3. Common Beginner Mistakes (and how to avoid them)

### Mistake 1: Expecting truthy / falsy values

This is not JavaScript or Python.

Wrong thinking:

```go
if age { // ❌
}
```

Correct:

```go
if age > 0 {
}
```

Rule:
**Only `true` or `false` is allowed. Nothing else.**

---

### Mistake 2: Overusing `else`

Beginners stack logic like this:

```go
if condition1 {
} else {
    if condition2 {
    } else {
        if condition3 {
        }
    }
}
```

This is unreadable trash.

Fix:

* Use `else if`
* Return early
* Or use `switch`

Clean code beats clever nesting.

---

### Mistake 3: Forgetting that `switch` exists

Beginners write long `if-else` chains for equality checks.

Bad:

```go
if status == "NEW" {
} else if status == "APPROVED" {
} else if status == "REJECTED" {
}
```

Better:

```go
switch status {
case "NEW":
case "APPROVED":
case "REJECTED":
}
```

If you ignore `switch`, you’re choosing worse code on purpose.

---

## 4. Real-World Scenarios Where Conditionals Matter

### Scenario 1: Business Rules & Validation (your daily reality)

Conditionals enforce:

* Currency checks
* Customer eligibility
* Order status transitions
* Error categorization

Every rule your company follows is encoded as a conditional somewhere.

---

### Scenario 2: Error Handling in Go

Go uses conditionals heavily for errors:

```go
if err != nil {
    return err
}
```

If you don’t respect this pattern, your Go code will be unidiomatic and fragile.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Takes an integer
* Prints whether it is positive, negative, or zero

---

### Exercise 2 (Medium)

Write a program that:

* Uses `switch`
* Determines shipping cost based on region code

---

### Exercise 3 (Hard)

Write a program that:

* Validates an order using multiple conditions
* Rejects invalid states early
* Avoids deep nesting

Design it cleanly.

---

## Thought-Provoking Question (Answer honestly)

Conditional statements **encode decisions**.

As systems grow, decision logic grows messy fast.

So here’s the real question:
When does adding **more conditionals** become a design smell, and what would you replace them with in a Go codebase?

If you can reason about that, you’re moving beyond beginner logic into real software design.
