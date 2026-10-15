# Testing and Benchmarking in Go

One of the best features of Go is that testing and benchmarking are built directly into the standard library. You do not need to download or install any external frameworks to ensure your code is reliable and fast.

## 1. Unit Testing

Tests check if your code works correctly and produces the expected outputs.

### The Two Golden Rules of Testing
1. **The File Name:** Test code must live in its own separate file. The file name must always end in `_test.go`. For example, if you are testing code in `main.go`, your test file should be named `main_test.go`.
2. **The Function Name:** Test functions must always start with the word `Test` followed by a capitalized name (e.g., `TestAdd`). Every test function must take exactly one parameter: `t *testing.T`.

### Example: Testing an Add Function

Imagine we have a simple function that adds two integers:
```go
func add(a, b int) int {
	return a + b
}
```

Here is how you would write a test for it:

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	// Call the function we want to test
	result := add(4, 7)
	
	// Check if the result matches our expectation
	if result != 11 {
		// If it fails, report the error using t.Errorf
		t.Errorf("Expected 11, got %d", result)
	}
}
```

**To run tests:** Open your terminal and simply type `go test`. Go will automatically find any file ending in `_test.go` and execute the tests inside.

## 2. Benchmarking

While tests check *if* your code works, benchmarks check *how fast* it runs. This is incredibly useful for comparing the performance of different functions.

### The Rules of Benchmarking
1. The function name must start with `Benchmark` instead of `Test` (e.g., `BenchmarkAdd`).
2. Instead of taking a `t *testing.T` parameter, it takes a `b *testing.B` parameter.

### How Benchmarks Work
Benchmarks do not just run your code once; they run it in a loop over and over again to calculate an average speed. Go dynamically provides a special number, `b.N`, which adjusts automatically until Go gets a statistically reliable measurement.

### Example: Benchmarking the Add Function

```go
package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	// Run the loop exactly b.N times
	for i := 0; i < b.N; i++ {
		add(4, 7)
	}
}
```

**To run benchmarks:** Open your terminal and type `go test -bench=.`. 

Go will output performance metrics that look something like this:
`BenchmarkAdd-8   1000000000   0.314 ns/op`

This tells you that Go ran your loop a billion times, and it took approximately 0.314 nanoseconds per operation!