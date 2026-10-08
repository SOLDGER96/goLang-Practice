package main

import (
	"fmt"
	"time"
)

type user struct {
	first string
	last string
	birth string
	create time.Time
}
func (u user) outputUserData() {
	fmt.Println(u.first, u.last, u.birth, u.create)
}

func main() {
	//
	firstName := getUserData("Enter your name: ")
	lastName := getUserData("Enter your Last name: ")
	birthDate := getUserData("Enter your Date of Birth (DD/MM/YY): ")
	
	var appUser user

	appUser = user{
		first: firstName,
		last: lastName,
		birth: birthDate,
		create: time.Now(),
	}
	// Here we are using the method attached to the struct we defined earlier
	appUser.outputUserData()
}

func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scan(&textInput)
	return textInput
}




