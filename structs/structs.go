package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *user) outputData() {
	output := fmt.Sprintf("First Name: %s || Last name: %s || Born in: %s", u.firstName, u.lastName, u.birthdate)
	fmt.Println(output)
}

func (u *user) clearUserData() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
}

// Creation/Constructor function for the user Struct
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate cannot be empty")
	}

	return &user{firstName: firstName, lastName: lastName, birthdate: birthdate, createdAt: time.Now()}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.outputData()
	appUser.clearUserData()
	appUser.outputData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// Uses the pointer to the user struct, to avoid copying the whole struct and also in GO it's not needed to dereference the pointer to access the struct fields
func outputUserData(newUser *user) {
	output := fmt.Sprintf("User: %s last name - %s -> was born at %s", newUser.firstName, newUser.lastName, newUser.birthdate)
	fmt.Println(output)
}
