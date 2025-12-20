package main

import "fmt"

func main() {
	age := 20
	agePointer := &age

	fmt.Println("Current age is ", age)

	fmt.Println("Age using pointer: ", agePointer)

	editAgeToAdultYears(&age)

	fmt.Println("Adult years: ", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
