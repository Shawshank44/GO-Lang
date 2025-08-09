package main

import (
	"math/rand"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// =======================TESTING===========================
//command : go test main_test.go

// Basic Test
// func TestAdd(t *testing.T) {
// 	result := Add(2, 4)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2,3) = %d, want %d", result, expected)
// 	}
// }

// Table Driven test
// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5}, //  (2 + 3 = 5) adding with the help of struct
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}

// 	for _, test := range tests {
// 		result := Add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("Add(%d,%d) = %d; want %d", test.a, test.b, result, test.expected)
// 		}
// 	}
// }

// func TestAddSubtests(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			result := Add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("result = %d; want := %d", result, test.expected)
// 			}
// 		})
// 	}
// }

// ================================BENCHMARKING========================================

// command : go test -bench="." main_test.go | Select-String -NotMatch "cpu:"
// command (for memory allocation) : go test -bench="." -benchmem  main_test.go | Select-String -NotMatch "cpu:"

// basic Benchmark
// func BenchmarkAdd(b *testing.B) {
// 	for range b.N {
// 		Add(2, 3)
// 	}
// }

// func BenchmarkAddSmallInput(b *testing.B) {
// 	for range b.N {
// 		Add(20, 30)
// 	}
// }

// func BenchmarkAddMediumInput(b *testing.B) {
// 	for range b.N {
// 		Add(200, 300)
// 	}
// }

// func BenchmarkAddLargeInput(b *testing.B) {
// 	for range b.N {
// 		Add(2000, 3000)
// 	}
// }

// ====================PROFILING========================
// command : go test -bench="." -memprofile mem.pprof  main_test.go | Select-String -NotMatch "cpu:"
// command (to execute mem.pprof): go tool pprof mem.pprof

func GeneratingRandomSlice(_size int) []int {
	slice := make([]int, _size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GeneratingRandomSlice(100)
	if len(slice) != size {
		t.Errorf("Expected slice size %d; received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GeneratingRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GeneratingRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		SumSlice(slice)
	}
}
