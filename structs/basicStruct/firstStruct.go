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

	// This format also works but be aware of the order of var
	// If values are skipped then they are assigned null values
	// appUser = user{
	// 	first: firstName,
	// 	last: lastName,
	// 	birth: birthDate,
	// 	create: time.Now(),
	// }

	outputUserData(appUser)
}

func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scan(&textInput)
	return textInput
}
// func outputUserData(f1, l1, b1 string){
// 	fmt.Println(f1, l1, b1)
// }

func outputUserData(u user) {
	fmt.Println(u.first, u.last, u.birth, u.create)
}


