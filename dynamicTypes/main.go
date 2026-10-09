package main

import "fmt"

func main() {
	a := 20
	b := 30
	add(a,b)
	op :=addNew(a,b)
	fmt.Print(op)
}
// this is used for anything
func add(a, b any) any {
	aInt, aok := a.(int)
	bInt, bok := b.(int)
	if aok && bok {
		return aInt + bInt
	}

	return nil
}

// Generic feature.
func addNew[T int | float64 | string](a, b T) T {
	return a + b
}
