package main

import "fmt"

func main() {
	numSlice := []int{1,2,3,4,5}
	sum := sumUp(1,10,15)
	fmt.Println(sum)

	// We can also pass a slice to variadic funcn
	nsum := sumUp(numSlice...)
	fmt.Println(nsum)
}

func sumUp(num ...int) int {
	sum := 1
	for _, val := range num {
		sum += val
	}
	return sum
}