package main

import "fmt"

func hobbyArray() {
	hobbies := [3]string{"reading", "gaming", "hiking"}

	fmt.Printf("My hobbies: %s \n\r", hobbies)

	fmt.Println("My first hobby is:", hobbies[0])
	fmt.Println("My second and third hobbies are:", hobbies[1:3])

	hobbySlice := hobbies[:2]
	//hobbySlice = hobbies[0:2]
	fmt.Printf("Hobby slice: %s \n\r", hobbySlice)

	finalSlice := hobbySlice[1:3]
	fmt.Printf("Final slice: %s \n\r", finalSlice)
}

func dynamicGoalArray() {
	goals := []string{"learn Go", "build projects"}

	goals[1] = "contribute to open source"
	goals = append(goals, "read Go books", "attend Go meetups")

	fmt.Printf("My goals: %s \n\r", goals)
}

func productArray() {
	products := []Product{
		{Id: 1, Price: 19.99},
		{Id: 2, Price: 29.99},
	}

	products = append(products, Product{Id: 1, Price: 19.99})

	fmt.Printf("Products: %v \n\r", products)
}

type Product struct {
	Id    int
	Price float64
}
