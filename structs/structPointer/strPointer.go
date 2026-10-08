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

	outputUserData(&appUser)
}

func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scan(&textInput)
	return textInput
}

func outputUserData(u *user) {
	fmt.Println(u.first, u.last, u.birth, u.create)
}


