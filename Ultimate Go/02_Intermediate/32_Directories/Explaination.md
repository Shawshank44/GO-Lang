Let’s cut through the fluff and tackle **directories** in Go—the containers for your files, and a core part of file system management.

---

## 1. Directories in Go — What They Are and Why They Exist

A **directory** (or folder) is a filesystem object that can contain files and other directories.

Purpose:

* Organize files logically
* Enable hierarchical storage and efficient access
* Support operations like listing, creating, moving, or deleting files in bulk

When commonly used:

* Managing configuration, assets, or logs
* Traversing directories for processing multiple files
* Dynamically creating project or temporary folders in programs

Truth check:

> Treat directories like first-class objects—you can’t just assume they exist.

---

## 2. Simple Code Example (Listing Directory Contents)

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "./exampleDir"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, f := range files {
		fmt.Println(f.Name(), "IsDir:", f.IsDir())
	}
}
```

Key points:

* `os.ReadDir` returns a slice of `DirEntry`
* `IsDir()` distinguishes between files and subdirectories
* Always handle errors; directories may not exist or be inaccessible

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Assuming a directory exists

* Trying to read or write files in a non-existent folder
* Causes runtime errors

Fix:

* Use `os.Mkdir` or `os.MkdirAll` to create directories safely before use

---

### Mistake 2: Not handling permissions

* Ignoring file system permissions
* Leads to “permission denied” errors

Fix:

* Check and handle permissions when creating directories
* Use appropriate mode bits (`os.FileMode`) for intended access

---

### Mistake 3: Confusing file vs directory operations

* Using file-specific functions on directories (e.g., `os.OpenFile`)
* Leads to errors or unexpected behavior

Fix:

* Use directory-aware functions like `os.ReadDir`, `os.Mkdir`, `filepath.WalkDir`

---

## 4. Real-World Applications

### Scenario 1: Log Management

* Organize logs by date, project, or module into directories
* Enable easier rotation, cleanup, and monitoring

---

### Scenario 2: File Processing Tools

* Traverse directories to process large datasets
* E.g., batch renaming, compression, or backup scripts

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Checks if a directory exists
* If it doesn’t, creates it using `os.MkdirAll`

---

### Exercise 2 (Medium)

Write a program that:

* Lists all files in a directory
* Separates files and subdirectories in the output

---

### Exercise 3 (Hard)

Write a program that:

* Recursively traverses a given directory
* Prints the full path of all `.go` files
* Skips hidden directories and handles permission errors gracefully

---

## Thought-Provoking Question

If your program needs to process user-specified directories, **how would you prevent malicious input from causing directory traversal attacks, infinite loops, or accidental deletion of critical system folders?**

Think about validation, sanitization, and safe traversal strategies.
