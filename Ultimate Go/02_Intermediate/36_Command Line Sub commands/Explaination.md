Here’s a detailed guide to **command-line subcommands** in Go:

---

## 1. Command-Line Subcommands — What They Are and Why They Exist

Subcommands are **commands within a CLI program** that perform distinct tasks.
Think of `git commit` or `git push`—`commit` and `push` are subcommands.

Purpose:

* Organize CLI programs with multiple functionalities
* Provide a clear, modular interface
* Separate flags and logic for each task

When commonly used:

* Multi-functional CLI tools
* Programs with logically distinct operations
* Scripts that manage multiple workflows

Truth check:

> Subcommands often have their own flags separate from global flags.

---

## 2. Simple Code Example

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'greet' or 'bye' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		name := greetCmd.String("name", "Guest", "Name to greet")
		greetCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *name)

	case "bye":
		byeCmd := flag.NewFlagSet("bye", flag.ExitOnError)
		name := byeCmd.String("name", "Guest", "Name to say goodbye")
		byeCmd.Parse(os.Args[2:])
		fmt.Printf("Goodbye, %s!\n", *name)

	default:
		fmt.Println("Expected 'greet' or 'bye' subcommands")
		os.Exit(1)
	}
}
```

Key points:

* Use `flag.NewFlagSet` for each subcommand
* Parse only the relevant arguments (`os.Args[2:]`)
* Each subcommand can have its own flags

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Parsing all flags together

* Trying to parse global flags before identifying the subcommand

Fix:

* Parse flags **after identifying the subcommand** using `flag.NewFlagSet`

---

### Mistake 2: Ignoring argument length checks

* Accessing `os.Args[1]` without checking if it exists can cause panics

Fix:

* Always check `len(os.Args)` before accessing subcommand arguments

---

### Mistake 3: Sharing flag pointers across subcommands

* Reusing pointers can cause unexpected behavior

Fix:

* Create **independent FlagSets** for each subcommand

---

## 4. Real-World Applications

### Scenario 1: Git-style tools

* Subcommands like `init`, `commit`, `push` make the CLI intuitive and organized

### Scenario 2: Multi-purpose system utilities

* A single program handling `backup`, `restore`, `status` as separate subcommands

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a CLI program with two subcommands:

* `hello` — prints "Hello, World!"
* `goodbye` — prints "Goodbye, World!"
  Each subcommand has a `-name` flag to customize the message.

---

### Exercise 2 (Medium)

Create a CLI program with three subcommands:

* `add` — takes two integer flags `-x` and `-y` and prints their sum
* `sub` — takes two integer flags and prints the difference
* `mul` — takes two integer flags and prints the product
  Handle missing or invalid flags gracefully.

---

### Exercise 3 (Hard)

Build a file management CLI with subcommands:

* `list` — lists files in a directory with optional `-ext` filter
* `copy` — copies a file from `-src` to `-dst`
* `delete` — deletes a file specified by `-file`
  Ensure proper error handling, argument validation, and help messages for each subcommand.

---

## Thought-Provoking Question

When designing a CLI with multiple subcommands and overlapping flags, **how would you structure your program to minimize confusion, avoid flag collisions, and provide a consistent user experience** while keeping the code maintainable?
