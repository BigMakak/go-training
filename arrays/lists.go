package main

import "fmt"

func main() {
	//dynamicArray()
	hobbyArray()

	dynamicGoalArray()

	productArray()
}

func dynamicArray() {
	// Dynamic array, no specified size
	values := []int{10, 20, 30}

	// Append values to the slice using append function
	values = append(values, 50, 60)

	// To remove values, create a new slice using the : operator
	values = append(values[2:]) // start on index 2 to the end and remove the first two values

	fmt.Println(values)
}

func arrayExercise() {

	// Semi initialized array of string
	names := [5]string{"product-1"}
	names[3] = "product-4"

	// Fully initialized array of float64
	prices := [4]float64{19.99, 29.99, 39.99, 49.99}

	// Slice of prices
	featuredPrices := prices[1:]

	// Slice of a slice
	highlightedPrices := featuredPrices[:1]

	// Slice capacity and length
	fmt.Printf("Length of featuredPrices: %d and capacity: %d \n\r", len(featuredPrices), cap(featuredPrices))
	fmt.Printf("length of highlighteddPrices: %d and capacity %d \n\r", len(highlightedPrices), cap(highlightedPrices))

	// Get more from the underlying array by extending the slice. There is alwasy more values to the right in the underlying array
	highlightedPrices = highlightedPrices[:3]
	fmt.Printf("Final length of highlighteddPrices: %d and capacity %d \r\n", len(highlightedPrices), cap(highlightedPrices))
	// Print highlighted prices
	fmt.Println(prices)
	fmt.Println(names)
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
}
