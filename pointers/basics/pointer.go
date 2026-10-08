package main

import "fmt"

func main() {
	age := 23

	agePointer := &age

	fmt.Println("Your age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Your Adult Years: ",adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
