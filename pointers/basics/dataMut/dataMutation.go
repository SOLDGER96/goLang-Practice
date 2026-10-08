package main

import "fmt"

func main() {
	age := 20

	fmt.Println(age)

	ageLoc := &age
	getTheAddress(ageLoc)

	fmt.Println(age)
}

func getTheAddress(age *int) {
	*age -= 10
}