package main

import (
	"example.com/structs/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)

	adminUser, errAdmin := user.NewAdmin("francis.gmail.com", "securepassword123")

	if errAdmin != nil {
		fmt.Println(errAdmin)
		return
	}

	adminUser.OutputAdminData()
	adminUser.ClearUserData()
	adminUser.OutputAdminData()

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.OutputData()
	appUser.ClearUserData()
	appUser.OutputData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
