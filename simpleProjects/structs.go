package main

import (
	"example.com/structs/user"
	"fmt"
	
)




func main() {
	//
	firstName := getUserData("Enter your name: ")
	lastName := getUserData("Enter your Last name: ")
	birthDate := getUserData("Enter your Date of Birth (DD/MM/YY): ")
	
	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Print(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "password")
	admin.OutputUserData()
	admin.User.ClearUserName()
	admin.User.OutputUserData()


	// Here we are using the method attached to the struct we defined earlier
	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData() 
}
func getUserData(textInput string) string {
	fmt.Print(textInput)
	fmt.Scanln(&textInput)
	return textInput
}





