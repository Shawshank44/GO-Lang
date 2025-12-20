Here’s a structured guide to **Logging in Go**:

---

## 1. Logging — What It Is and Why It Exists

Logging is the process of **recording runtime information about a program**, such as events, errors, or debug information.

Purpose:

* Monitor program execution and diagnose issues
* Track errors and system behavior for debugging
* Maintain audit trails for operations

Commonly used:

* Debugging during development
* Monitoring production systems
* Recording important system events or errors

---

## 2. Simple Code Example

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Set output to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("Application started")
	log.Println("Performing some operations...")

	err = doTask()
	if err != nil {
		log.Printf("Error occurred: %v", err)
	}
}

func doTask() error {
	// Simulate an error
	return nil
}
```

Key points:

* `log.Println` prints info with timestamp
* `log.Printf` allows formatted logging
* `log.SetOutput` can redirect logs to files or other writers

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Logging sensitive information

* Avoid printing passwords, tokens, or personal user data

Fix:

* Mask or omit sensitive info before logging

---

### Mistake 2: Not configuring output destination

* Logs going only to stdout may be lost in production

Fix:

* Configure logs to write to files, syslog, or external logging systems

---

### Mistake 3: Overusing or underusing logs

* Too many logs create noise; too few make debugging difficult

Fix:

* Use appropriate log levels (info, warning, error)
* Log only meaningful events

---

## 4. Real-World Applications

### Scenario 1: Production monitoring

* Track errors, warnings, and system health in a live server

### Scenario 2: Audit trails

* Record critical operations like user authentication or financial transactions

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a Go program that logs the start and end of the program execution to a file named `program.log`.

---

### Exercise 2 (Medium)

Write a Go program that logs user login attempts. Include timestamps and whether the attempt was successful or failed.

---

### Exercise 3 (Hard)

Create a Go program that:

1. Logs messages at different levels (info, warning, error)
2. Redirects info logs to one file and error logs to another
3. Demonstrates logging an error with stack trace or formatted details

---

## Thought-Provoking Question

If logging is crucial for debugging and monitoring, **how would you design a logging strategy for a high-load production system to ensure performance isn’t degraded while still capturing useful diagnostic information**?
