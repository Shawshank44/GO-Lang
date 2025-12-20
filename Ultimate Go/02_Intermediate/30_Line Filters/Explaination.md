Alright, let’s tackle **line filters** in Go. These are deceptively simple but extremely powerful for processing streams of text.

---

## 1. Line Filters in Go — What They Are and Why They Exist

**Line filters** process text **line by line**, usually reading input, transforming or filtering lines, and writing output.

Purpose:

* Process large text or log files efficiently
* Avoid loading the entire file into memory
* Apply transformations, searches, or selective output

When it’s commonly used:

* Log analysis (filtering out errors or warnings)
* Data pipelines (transforming CSV/JSON lines)
* Command-line utilities (like Unix `grep`, `sed`)

Hard truth:

> Reading everything into memory for “filtering lines” is inefficient and unnecessary for large datasets.

---

## 2. Simple Code Example (Filter Lines Containing a Word)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
```

Key points:

* Uses `bufio.Scanner` for line-by-line reading
* Efficient for large files
* Filters lines containing `"ERROR"`

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Loading entire file into memory

* Beginners often use `os.ReadFile` to filter lines
* Crashes or high memory usage on large files

Fix:

* Use `bufio.Scanner` or streaming approaches for line-by-line processing

---

### Mistake 2: Forgetting to handle errors

* Skipping `scanner.Err()` or `file.Open()` errors
* Can silently fail or produce incomplete output

Fix:

* Always check errors after opening files and after scanning

---

### Mistake 3: Mismanaging line endings

* Confusing `\n` vs `\r\n` on different OSes
* Filters fail or output is incorrect

Fix:

* Use `scanner.Text()` instead of manual splitting
* Consider normalizing line endings if needed

---

## 4. Real-World Applications

### Scenario 1: Log Monitoring

* Extract only relevant lines (errors, warnings, failed transactions)
* Feed filtered logs into monitoring or alerting systems

---

### Scenario 2: Data Cleaning Pipelines

* Remove or transform lines from CSV, JSONL, or TSV files
* Preprocess input before further analytics or ingestion

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Reads a text file line by line
* Prints only lines that contain the word `"Go"`

---

### Exercise 2 (Medium)

Write a program that:

* Reads a CSV file line by line
* Prints lines where the second column contains a number greater than 100

---

### Exercise 3 (Hard)

Write a program that:

* Reads a log file line by line
* Filters lines containing `"ERROR"` or `"WARN"`
* Adds a timestamp prefix to each filtered line before printing
* Ensures memory usage is minimal for very large files

---

## Thought-Provoking Question

If you process a file line by line using a line filter, **how would you handle a scenario where one line depends on the content of the previous line?**

Think about the trade-off between memory efficiency and context-aware processing.
