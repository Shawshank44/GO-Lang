Good choice. `bufio` is one of those packages people *use blindly* without understanding why it exists, and that’s how performance bugs and subtle I/O issues creep in. Let’s do this properly.

---

## 1. What `bufio` Is (Purpose & When It’s Used)

`bufio` provides **buffered I/O** for reading from and writing to streams like files, network connections, or standard input/output.

Purpose:

* Reduce the number of expensive system calls
* Read input **incrementally** (line by line, word by word, etc.)
* Write output efficiently instead of flushing every byte

Commonly used when:

* Reading large files
* Processing user input from `stdin`
* Handling network connections
* Writing logs or streamed output

Hard truth:

> If you read or write frequently without buffering, you are wasting CPU time on syscalls.

---

## 2. Simple Code Example

### Reading input using `bufio.Reader`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Hello,", name)
}
```

What’s happening:

* `bufio.NewReader` wraps `os.Stdin`
* Input is read **buffered**, not byte-by-byte
* `ReadString` stops at a delimiter

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Thinking `bufio` is optional for performance

It’s not.

Wrong assumption:

> “The OS will handle it anyway.”

Reality:

* Each unbuffered read/write can trigger a syscall
* Thousands of syscalls = slow program

Fix:

* Use `bufio.Reader` and `bufio.Writer` whenever doing repeated I/O

---

### Mistake 2: Forgetting to flush `bufio.Writer`

This is classic.

Example bug:

```go
writer := bufio.NewWriter(os.Stdout)
writer.WriteString("Hello")
// Output never appears
```

Why:

* Data stays in memory

Fix:

```go
writer.Flush()
```

Rule:

> If you use `bufio.Writer`, flushing is **your responsibility**.

---

### Mistake 3: Misusing `Scanner` for large input

`bufio.Scanner` has a **token size limit** (64K by default).

Problem:

* Large lines silently fail or error

Fix:

* Use `bufio.Reader` for large or unbounded input
* Or manually increase scanner buffer size if appropriate

---

## 4. Real-World Applications Where `bufio` Is Useful

### Scenario 1: File Processing Tools

* Reading log files line by line
* Parsing CSVs or text-based data
* Streaming large datasets without loading everything into memory

Without buffering:

* Your program will crawl
* Memory usage spikes unnecessarily

---

### Scenario 2: Network Services

* Reading requests from TCP connections
* Writing responses efficiently
* Streaming protocols (chat servers, proxies)

This is **non-negotiable** for scalable systems.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Reads multiple lines from standard input
* Prints each line with a line number
* Stops when the user enters an empty line

---

### Exercise 2 (Medium)

Create a file reader that:

* Reads a text file line by line
* Counts total lines and total characters
* Uses buffered reading throughout

---

### Exercise 3 (Hard)

Build a buffered logger that:

* Writes log messages to a file
* Flushes periodically instead of after every write
* Handles write errors correctly

This will expose whether you actually understand buffering or not.

---

## Thought-Provoking Question

If buffering improves performance but delays I/O visibility, **how do you decide when correctness matters more than speed—and who gets to make that decision: the library, or your application logic?**

Think carefully. This tradeoff shows up everywhere in systems programming.
