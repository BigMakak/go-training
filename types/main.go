package main

import "fmt"

func main() {
	printSomething("hello")
	printSomething(1)
	printSomething(6.4)

	sumInt := addGeneric(3, 5)
	sumStrings := addGeneric("Hello, ", "World!")
	sumFloat := addGeneric(4.5, 3.2)

	fmt.Println(sumInt)
	fmt.Println(sumStrings)
	fmt.Println(sumFloat)
}

// Print different types of data using different ways of checking the types
func printSomething(data interface{}) {
	switch v := data.(type) {
	case string:
		fmt.Println("String data:", v)
	case int:
		fmt.Println("Integer data:", v)
	case float64:
		fmt.Println("Float data:", v)
	default:
		fmt.Println("Unknown type")
	}

	intValue, ok := data.(int)
	if ok {
		fmt.Println("The integer value is:", intValue)
	}

	floatValue, ok := data.(float64)
	if ok {
		fmt.Println("The float value is:", floatValue)
	}

	stringValue, ok := data.(string)
	if ok {
		fmt.Println("The string value is:", stringValue)
	}
}

// Generic function to add two values that are from a generic type/group
func addGeneric[T int | float64 | string](a, b T) T {
	return a + b
}
