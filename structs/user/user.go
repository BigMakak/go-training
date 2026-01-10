package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputData() {
	output := fmt.Sprintf("First Name: %s || Last name: %s || Born in: %s", u.firstName, u.lastName, u.birthdate)
	fmt.Println(output)
}

func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}

func (a *Admin) OutputAdminData() {
	a.OutputData()
	output := fmt.Sprintf("Admin Email: %s || Password: %s", a.email, a.password)
	fmt.Println(output)
}

// New Creation/Constructor function for the user Struct
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate cannot be empty")
	}

	return &User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password cannot be empty")
	}

	return &Admin{
		email:    email,
		password: password,
		User:     User{"Admin", "User", "01/01/1970", time.Now()},
	}, nil
}
