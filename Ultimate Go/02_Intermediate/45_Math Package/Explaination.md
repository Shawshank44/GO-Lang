Here’s a structured guide to the **`math` package in Go**:

---

## 1. `math` Package — What It Is and Why It Exists

The `math` package provides **basic constants and mathematical functions** for floating-point operations.

Purpose:

* Perform common mathematical calculations like power, square root, logarithms, trigonometry, rounding, etc.
* Simplify computations that would otherwise require manual implementation
* Ensure precision and efficiency for floating-point operations

Commonly used:

* Scientific and engineering computations
* Geometry calculations
* Financial or statistical calculations requiring precise mathematical operations

---

## 2. Simple Code Example

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 16.0
	y := 3.0

	fmt.Println("Square root of x:", math.Sqrt(x))
	fmt.Println("x raised to power y:", math.Pow(x, y))
	fmt.Println("Sine of 45 degrees:", math.Sin(45*math.Pi/180))
	fmt.Println("Logarithm of x:", math.Log(x))
	fmt.Println("Ceil of 2.3:", math.Ceil(2.3))
	fmt.Println("Floor of 2.7:", math.Floor(2.7))
}
```

Key points:

* Most functions operate on `float64` and return `float64`
* Trigonometric functions expect radians
* Constants like `math.Pi` provide precision

---

## 3. Common Beginner Mistakes (and How to Avoid Them)

### Mistake 1: Confusing degrees with radians

* Trigonometric functions (`Sin`, `Cos`, `Tan`) expect radians, not degrees

Fix:

* Convert degrees to radians: `radians = degrees * math.Pi / 180`

---

### Mistake 2: Using integer types

* Many `math` functions only accept `float64` and will cause type errors with `int`

Fix:

* Convert integers to `float64` when calling math functions: `math.Sqrt(float64(n))`

---

### Mistake 3: Ignoring floating-point precision issues

* Direct comparison of floating-point results may lead to unexpected behavior

Fix:

* Compare using a small epsilon: `math.Abs(a-b) < 1e-9`

---

## 4. Real-World Applications

### Scenario 1: Scientific computation

* Calculating distances, forces, angles, or other physics/engineering formulas

### Scenario 2: Financial calculations

* Compound interest, logarithmic growth, or rounding calculations in banking software

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Calculate the hypotenuse of a right-angled triangle given sides `a` and `b` using `math.Sqrt` and `math.Pow`.

---

### Exercise 2 (Medium)

Write a function that takes an angle in degrees and returns its sine, cosine, and tangent values.

---

### Exercise 3 (Hard)

Create a small program to calculate the compound interest on an investment for `n` years, using `math.Pow`, and round the final amount to 2 decimal places.

---

## Thought-Provoking Question

Since the `math` package functions work with `float64`, how does the reliance on floating-point arithmetic impact precision in financial or scientific computations, and what strategies could you use in Go to mitigate potential rounding errors?
