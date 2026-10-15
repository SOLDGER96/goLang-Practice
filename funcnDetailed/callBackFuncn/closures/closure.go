package main

import "fmt"

func newCounter() func() int {
	count := 0
	return func() int {
		count ++
		return count
	}
}

func main() {
	c := newCounter()
	fmt.Println(c())
	fmt.Println(c())

	cb := newCounter()
	fmt.Println(cb())
	fmt.Println(cb())
}