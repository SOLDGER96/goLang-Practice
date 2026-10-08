package main

import (
	"errors"
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
// Make sure to ref to a pointeter when modifying structs or else it 
// will modify only the copy of that struct not the original struct
func (u *user) clearUserName() {
	u.first = ""
	u.last = ""
}
// Create a helper funcn to create a struct (constructer funcn)
func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Fields cannot be empty")
	}
	return &user{
		first: firstName,
		last: lastName,
		birth: birthDate,
		create: time.Now(),
	}, nil
}


func main() {
	//
	firstName := getUserData("Enter your name: ")
	lastName := getUserData("Enter your Last name: ")
	birthDate := getUserData("Enter your Date of Birth (DD/MM/YY): ")
	
	var appUser *user


	// appUser = user{
	// 	first: firstName,
	// 	last: lastName,
	// 	birth: birthDate,
	// 	create: time.Now(),
	// }

	// newUser Function can be used to create a struct avoiding the above multiple lines with helper funcn
	appUser, err := newUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Print(err)
		return
	}

	// Here we are using the method attached to the struct we defined earlier
	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
	fmt.Print(time.Now())
}

func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scanln(&textInput)
	return textInput
}




