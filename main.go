package main

import "fmt"

// write a function to add two numbers oinly using two varibles both function
// will not return any values just print the adddition in main function only

func main() {
	a, b := 3, 4
	adder(&a, &b)
	fmt.Println(a)
}
func adder(a, b *int) {
	*a += *b
}
