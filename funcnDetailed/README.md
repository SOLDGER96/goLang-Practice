# Callbacks and Closures in Go

In Go, functions are considered **"first-class citizens."** This means you can treat a function just like any other variable—you can assign it to a variable, or you can pass it into another function as an argument.

---

## 1. Callback Functions

When you pass a function into another function to be executed later, that passed function is called a **Callback**.

If we want to pass a simple function that takes no arguments and returns nothing, its type is written as `func()`. When passing the function as an argument, you just use its name without the parentheses.

### Example: A Simple Callback

```go
package main

import "fmt"

// ExecuteTask accepts a callback function as a parameter
func ExecuteTask(callBack func()) {
	// Trigger the passed-in function
	callBack() 
}

// Our standalone function that we will use as a callback
func callBack() {
	fmt.Println("Just a callback function!")
}

func main() {
	// Pass the function by name (no parentheses)
	ExecuteTask(callBack)
}
```

---

## 2. Closures

Closures go hand-in-hand with callbacks. A **closure** is simply an anonymous function that remembers the variables around it, even after the outer function that created those variables has finished running. It "closes over" or traps those variables in memory safely.

### Example: A Custom Counter using Closures

To see this magic in action, we can build a custom counter by writing a function that returns another function.

```go
package main

import "fmt"

// newCounter returns an anonymous function of type func() int
func newCounter() func() int {
	count := 0 // This is the variable our inner function will trap!
	
	// Return the anonymous function
	return func() int {
		count++      // Increase the trapped count by 1
		return count // Hand the new number back
	}
}

func main() {
	// counterA traps its own 'count' variable starting at 0
	counterA := newCounter()
	
	fmt.Println(counterA()) // Outputs: 1
	fmt.Println(counterA()) // Outputs: 2 (Remembers the previous state!)
	
	// counterB gets a brand new closure, trapping a completely separate 'count'
	counterB := newCounter()
	fmt.Println(counterB()) // Outputs: 1
}
```

### Why Closures are Powerful
Normally, when a function finishes running, all of its local variables are destroyed. But because the inner anonymous function in `newCounter` uses the `count` variable, Go keeps `count` alive safely in memory, trapped inside that specific function. This is incredibly useful for creating functions that need to remember state without relying on messy global variables.