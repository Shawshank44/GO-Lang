Good. Input/output (I/O) is where beginner Go programs either **look clean or fall apart fast**. I’ll keep this tight, practical, and honest.

---

## 1. What Output Functions and Input Functions Are

### Output Functions

Output functions **display data to the user or write it somewhere**.

Common ones:

* `fmt.Print`, `fmt.Println`, `fmt.Printf`
* Writing to files or streams via `os.Stdout`

Purpose:

* Communicate program state
* Show results
* Log information

Used when:

* Building CLI tools
* Debugging
* Reporting results

---

### Input Functions

Input functions **read data from the user or another source**.

Common ones:

* `fmt.Scan`, `fmt.Scanln`, `fmt.Scanf`
* `bufio.Reader` for line-based or large input

Purpose:

* Accept user input
* Read configuration or commands
* Process streamed data

Used when:

* Interactive programs
* CLI utilities
* Parsing structured input

---

## 2. Simple Code Example

```go
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

Key points:

* Input needs **addresses (`&`)**
* Output formatting matters
* This is fine for simple programs—not for robust input handling

---

## 3. Common Beginner Mistakes (and how to avoid them)

### Mistake 1: Forgetting to pass pointers to input functions

This is extremely common.

Wrong:

```go
fmt.Scan(name)
```

Fix:

* Always pass variable addresses

```go
fmt.Scan(&name)
```

---

### Mistake 2: Assuming `fmt.Scan` reads entire lines

It doesn’t.

`Scan` splits on whitespace.

Fix:

* Use `bufio.Reader` for full-line input
* Know what level of input control you need

---

### Mistake 3: Ignoring input errors

Beginners rarely check return values.

That’s reckless.

Fix:

* Always check the returned error
* Invalid input is normal, not exceptional

---

## 4. Real-World Scenarios Where I/O Is Central

### Scenario 1: Command-Line Tools

* Read arguments or user input
* Print structured output
* Exit with meaningful messages

Most Go tools live and die by clean I/O.

---

### Scenario 2: Data Processing Pipelines

* Read input from STDIN
* Process it
* Write results to STDOUT

This is how Go fits into Unix-style pipelines.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Takes two integers from the user
* Prints their sum using formatted output

---

### Exercise 2 (Medium)

Write a program that:

* Reads a full line of text from the user
* Prints the length of the input
* Handles spaces correctly

---

### Exercise 3 (Hard)

Build a CLI program that:

* Continuously reads user commands
* Prints different outputs based on input
* Handles invalid input without crashing

Focus on robustness, not speed.

---

## Thought-Provoking Question (Think like an engineer)

User input is **untrusted by default**.

So answer this:
**At what point should a Go program stop using `fmt.Scan` and move to buffered or structured input, and what risks do you introduce by sticking with simple I/O in larger systems?**

If you can reason about that trade-off, you’re past beginner level.
