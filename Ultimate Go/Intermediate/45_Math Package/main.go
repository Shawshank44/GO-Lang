package main

import (
	"fmt"
	"math"
)

func main() {
	// --- Constants ---
	fmt.Println("Constants:")
	fmt.Println("Pi:", math.Pi)
	fmt.Println("E:", math.E)
	fmt.Println("Phi:", math.Phi)
	fmt.Println("Sqrt2:", math.Sqrt2)
	fmt.Println("SqrtE:", math.SqrtE)
	fmt.Println("SqrtPi:", math.SqrtPi)
	fmt.Println("SqrtPhi:", math.SqrtPhi)
	fmt.Println("Ln2:", math.Ln2)
	fmt.Println("Log2E:", math.Log2E)
	fmt.Println("Ln10:", math.Ln10)
	fmt.Println("Log10E:", math.Log10E)
	fmt.Println()

	// --- Special Floating-Point Values ---
	fmt.Println("Special Floating-Point Values:")
	fmt.Println("Inf(+1):", math.Inf(+1))
	fmt.Println("Inf(-1):", math.Inf(-1))
	fmt.Println("IsInf(Inf, 1):", math.IsInf(math.Inf(1), 1))
	fmt.Println("NaN():", math.NaN())
	fmt.Println("IsNaN(NaN):", math.IsNaN(math.NaN()))
	fmt.Println()

	// --- Basic Arithmetic ---
	fmt.Println("Basic Arithmetic:")
	fmt.Println("Abs(-5.2):", math.Abs(-5.2))
	fmt.Println("Max(2.5, 3.5):", math.Max(2.5, 3.5))
	fmt.Println("Min(2.5, 3.5):", math.Min(2.5, 3.5))
	fmt.Println("Mod(5.5, 2.0):", math.Mod(5.5, 2.0))
	fmt.Println("Remainder(5.5, 2.0):", math.Remainder(5.5, 2.0))
	fmt.Println("Copysign(-3.0, 2.0):", math.Copysign(-3.0, 2.0))
	fmt.Println()

	// --- Rounding ---
	fmt.Println("Rounding:")
	fmt.Println("Ceil(2.3):", math.Ceil(2.3))
	fmt.Println("Floor(2.7):", math.Floor(2.7))
	fmt.Println("Round(2.5):", math.Round(2.5))
	fmt.Println("Trunc(2.7):", math.Trunc(2.7))
	fmt.Println("RoundToEven(3.5):", math.RoundToEven(3.5))
	fmt.Println()

	// --- Power & Exponentials ---
	fmt.Println("Power & Exponentials:")
	fmt.Println("Pow(2, 3):", math.Pow(2, 3))
	fmt.Println("Exp(1):", math.Exp(1))
	fmt.Println("Exp2(3):", math.Exp2(3))
	fmt.Println("Expm1(1):", math.Expm1(1)) // e^x - 1
	fmt.Println("Log(10):", math.Log(10))
	fmt.Println("Log2(8):", math.Log2(8))
	fmt.Println("Log10(100):", math.Log10(100))
	fmt.Println("Log1p(1):", math.Log1p(1)) // log(1+x)
	fmt.Println()

	// --- Roots & Hypotenuse ---
	fmt.Println("Roots & Hypotenuse:")
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("Cbrt(27):", math.Cbrt(27))
	fmt.Println("Hypot(3, 4):", math.Hypot(3, 4)) // sqrt(3^2 + 4^2)
	fmt.Println()

	// --- Trigonometry ---
	fmt.Println("Trigonometry:")
	fmt.Println("Sin(Pi/2):", math.Sin(math.Pi/2))
	fmt.Println("Cos(0):", math.Cos(0))
	fmt.Println("Tan(Pi/4):", math.Tan(math.Pi/4))
	fmt.Println("Asin(1):", math.Asin(1))
	fmt.Println("Acos(1):", math.Acos(1))
	fmt.Println("Atan(1):", math.Atan(1))
	fmt.Println("Atan2(1, 1):", math.Atan2(1, 1))
	fmt.Println()

	// --- Hyperbolic Functions ---
	fmt.Println("Hyperbolic Functions:")
	fmt.Println("Sinh(1):", math.Sinh(1))
	fmt.Println("Cosh(1):", math.Cosh(1))
	fmt.Println("Tanh(1):", math.Tanh(1))
	fmt.Println("Asinh(1):", math.Asinh(1))
	fmt.Println("Acosh(2):", math.Acosh(2))     // Valid only for x >= 1
	fmt.Println("Atanh(0.5):", math.Atanh(0.5)) // Valid only for -1 < x < 1
	fmt.Println()

	// --- Floating Point Helpers ---
	fmt.Println("Floating Point Helpers:")
	frac, exp := math.Frexp(8) // 8 = frac × 2^exp
	fmt.Println("Frexp(8):", frac, exp)
	fmt.Println("Ldexp(0.5, 4):", math.Ldexp(0.5, 4))
	fmt.Println("Nextafter(1, 2):", math.Nextafter(1, 2))
	intPart, fracPart := math.Modf(3.14)
	fmt.Println("Modf(3.14): int =", intPart, ", frac =", fracPart)
	fmt.Println("Signbit(-5):", math.Signbit(-5))
	fmt.Println("Dim(5, 3):", math.Dim(5, 3))
	fmt.Println()

	// --- Gamma & Special Functions ---
	fmt.Println("Gamma & LogGamma:")
	fmt.Println("Gamma(5):", math.Gamma(5)) // Γ(5) = 4! = 24
}
