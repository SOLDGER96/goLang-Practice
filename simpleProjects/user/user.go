package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	first string
	last string
	birth string
	create time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			first: "ADMIN",
			last: "ADMIN",
			birth: "-------",
			create: time.Now(),
		},
	}
}

func (u User) OutputUserData() {
	fmt.Println(u.first, u.last, u.birth, u.create)
}
// Make sure to ref to a pointeter when modifying structs or else it 
// will modify only the copy of that struct not the original struct
func (u *User) ClearUserName() {
	u.first = ""
	u.last = ""
}
// Create a helper funcn to create a struct (constructer funcn)
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Fields cannot be empty")
	}
	return &User{
		first: firstName,
		last: lastName,
		birth: birthDate,
		create: time.Now(),
	}, nil
}
