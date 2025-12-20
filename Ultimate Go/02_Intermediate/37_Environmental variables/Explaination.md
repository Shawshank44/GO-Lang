Here’s a structured guide to **Environment Variables in Go**:

---

## 1. Environment Variables — What They Are and Why They Exist

Environment variables are **key-value pairs set outside a program** that the program can read at runtime.

Purpose:

* Store configuration outside code (e.g., API keys, database URLs)
* Avoid hardcoding sensitive or environment-specific values
* Make programs portable across systems without changing code

Commonly used:

* Configuring credentials, paths, or modes (dev/prod)
* Feature toggles or flags that depend on deployment environment
* Any setting that may change between machines or deployments

---

## 2. Simple Code Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Read environment variable
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "defaultUser"
	}
	fmt.Println("Database User:", dbUser)

	// Set environment variable (for current process)
	os.Setenv("APP_MODE", "development")
	fmt.Println("App Mode:", os.Getenv("APP_MODE"))
}
```

Key points:

* `os.Getenv(key)` reads a variable
* `os.Setenv(key, value)` sets a variable for the current process
* Always check for empty values to avoid unintended behavior

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Assuming the variable always exists

* `os.Getenv` returns an empty string if the variable is unset

Fix:

* Always check for empty string and provide a default or handle the error

---

### Mistake 2: Thinking `os.Setenv` changes system environment

* `os.Setenv` only affects the current process, not the global system environment

Fix:

* Use configuration files or external scripts for persistent system-level environment variables

---

### Mistake 3: Storing sensitive information insecurely

* Hardcoding secrets into code instead of using environment variables

Fix:

* Keep secrets in environment variables or secret managers; never commit them in code

---

## 4. Real-World Applications

### Scenario 1: Database configuration

* `DB_HOST`, `DB_USER`, `DB_PASSWORD` can be configured differently for dev, staging, and production

### Scenario 2: Feature toggles

* Use an environment variable like `FEATURE_X_ENABLED=true` to enable or disable features without code changes

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a Go program that reads an environment variable `USERNAME` and prints `"Hello, <USERNAME>!"`. If not set, print `"Hello, Guest!"`.

---

### Exercise 2 (Medium)

Write a Go program that reads multiple environment variables: `APP_ENV` (dev/prod) and `LOG_LEVEL` (info/debug). Print appropriate messages depending on their values. Provide defaults if not set.

---

### Exercise 3 (Hard)

Create a Go program that:

1. Reads database configuration (`DB_HOST`, `DB_USER`, `DB_PASS`) from environment variables
2. Validates that none of them are empty, otherwise exits with an error
3. Prints a connection string formatted as: `user=<DB_USER> password=<DB_PASS> host=<DB_HOST>`

---

## Thought-Provoking Question

If environment variables are widely used for configuration and secrets, **what strategies would you adopt to prevent accidental exposure of sensitive environment variables while maintaining portability and ease of deployment**?
