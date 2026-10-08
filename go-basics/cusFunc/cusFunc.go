package main

import "fmt"



func main() {
	opText("Tanmay")

	// Custom function for adding and subtraction
	n1 := 30
	n2 := 20
	addition, subtraction := addSub(n1, n2)
	fmt.Printf("Addition: %v\nSubtraction: %v\n", addition, subtraction)
}

func opText(t string) {
	fmt.Println(t)
}
func addSub(a, b int) (int, int) {
	add := a+b
	sub := a-b
	return add, sub
}