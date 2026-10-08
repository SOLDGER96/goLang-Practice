package main

import (
	"fmt"
	//"time"
)

// type user struct {
// 	first string
// 	last string
// 	birth string
// 	create time.Time
// }

func main() {
	//
	firstName := getUserData("Enter your name: ")
	lastName := getUserData("Enter your Last name: ")
	birthDate := getUserData("Enter your Date of Birth (DD/MM/YY): ")
	
	outputUserData(firstName, lastName, birthDate)
}

func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scan(&textInput)
	return textInput
}
func outputUserData(f1, l1, b1 string){
	fmt.Println(f1, l1, b1)
}