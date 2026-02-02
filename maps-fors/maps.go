package main

import "fmt"

// Custom types
// Used to make types more readable for developers
type stringMap map[string]string

// We can add custom functions to these types
func (s *stringMap) print() {
	for key, value := range *s {
		fmt.Println("Key:", key, " || Value:", value)
	}
}

func main() {
	//mapExercise()
	makeCollections()
}

func mapExercise() {
	websites := map[string]string{
		"Google":        "https://www.google.com",
		"GitHub":        "https://www.github.com",
		"StackOverflow": "https://stackoverflow.com",
	}

	fmt.Println(websites)

	// Print google website URL
	fmt.Println(websites["Google"])

	// Add a new website
	websites["Reddit"] = "https://www.reddit.com"

	// Update an existing website
	websites["GitHub"] = "https://github.com"

	fmt.Println(websites)
}

// make used for creating slices and maps for a better memory allocation/performance
func makeCollections() {
	// Create a slice integers instead of a fixed array
	// the make function allocates memory for the slice, with an initial length and max capacity
	// With this slice when adding more elements it will be more efficient as it does not need to reallocate memory each time
	// this still can grow beyond the initial capacity, but is not that efficient
	userAges := make([]int, 1, 3) // length 1, capacity 3

	userAges = append(userAges, 25)
	userAges = append(userAges, 30)
	userAges = append(userAges, 22) // go will reallocate memory here due to passing the original capacity

	for index, age := range userAges {
		fmt.Printf("User Index %d age: %d \n\r", index, age)
	}

	// Create a map using make function
	userPhones := make(stringMap, 2)

	userPhones["Alice"] = "123-456-7890"
	userPhones["Bob"] = "987-654-3210"
	userPhones["Charlie"] = "555-555-5555" // again, go will reallocate memory here

	userPhones.print()
}
