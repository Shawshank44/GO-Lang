Here’s a structured guide to **JSON in Go**:

---

## 1. JSON — What It Is and Why It Exists

**JSON (JavaScript Object Notation)** is a lightweight, text-based format for **storing and exchanging structured data**.

Purpose:

* Facilitate communication between systems (e.g., APIs)
* Store configuration or structured data in a readable format
* Serialize and deserialize Go structs for persistence or transmission

Commonly used:

* Sending/receiving data in web applications
* Configuration files
* API request and response bodies

---

## 2. Simple Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	user := User{Name: "Alice", Email: "alice@example.com", Age: 30}

	// Convert struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", string(jsonData))

	// Convert JSON back to struct
	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded Struct: %+v\n", decodedUser)
}
```

Key points:

* `json.Marshal` serializes Go struct → JSON
* `json.Unmarshal` deserializes JSON → Go struct
* Use struct tags (`json:"fieldname"`) to control JSON keys

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Ignoring struct tags

* Without `json:"name"`, field names in JSON use Go struct field names (capitalized)

Fix:

* Always define JSON tags for correct key names

---

### Mistake 2: Using non-exported fields

* Fields starting with lowercase letters are **not exported** and cannot be marshaled/unmarshaled

Fix:

* Always use uppercase first letter for fields you want in JSON

---

### Mistake 3: Not handling errors

* Ignoring errors from `Marshal` or `Unmarshal` can hide parsing issues

Fix:

* Always check for errors when serializing/deserializing

---

## 4. Real-World Applications

### Scenario 1: API communication

* Web services exchanging JSON between client and server

### Scenario 2: Configuration files

* Storing app settings in JSON for readability and easy parsing

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Create a Go struct representing a `Book` with `Title`, `Author`, and `Price`. Serialize it to JSON and print the result.

---

### Exercise 2 (Medium)

Write a Go program that reads a JSON string containing a list of `User` objects and prints the names of users older than 25.

---

### Exercise 3 (Hard)

Create a Go program that:

1. Reads a JSON file containing multiple orders (each with `ID`, `Customer`, `Items`, `Total`)
2. Unmarshals the JSON into structs
3. Calculates and prints the total revenue
4. Marshals a filtered list of orders (e.g., orders > $100) back to JSON and writes it to a new file

---

## Thought-Provoking Question

When designing APIs that exchange JSON data, **how would you handle versioning and backward compatibility of JSON structures to ensure clients don’t break when fields are added, removed, or renamed**?
