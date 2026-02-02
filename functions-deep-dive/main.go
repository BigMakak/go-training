package main

import "fmt"

// Define a function type to be easily reused as an argument type for functions
type numberTransform func(int) int
type functionType int

const (
	Double functionType = iota
	Triple
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubledNumbers := transformNumbers(&numbers, transformerFunction(Double))
	tripledNumbers := transformNumbers(&numbers, transformerFunction(Triple))
	anonymousNumbers := transformNumbers(&numbers, func(number int) int {
		return number + 10
	})

	closerNumbers := transformNumbers(&numbers, transformerCreator(4))
	factorialNumber := transformNumbers(&numbers, recursiveFactorial)
	variadicSum := sumNumbers(1, 2, 3, 4, 5, 6)
	variadicSumArray := sumNumbers(numbers...) // spread all the elements of the array/slice to be passed as variadic parameters

	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Doubled Numbers:", doubledNumbers)
	fmt.Println("Tripled Numbers:", tripledNumbers)
	fmt.Println("Anonymous Transformed Numbers :", anonymousNumbers)
	fmt.Println("Closer Numbers:", closerNumbers)
	fmt.Println("Factorial Numbers:", factorialNumber)
	fmt.Println("Variadic Sum:", variadicSum)
	fmt.Println("Variadic Sum from Array:", variadicSumArray)
}

// Using function types as parameters
func transformNumbers(numbers *[]int, transform numberTransform) []int {
	transformed := []int{}

	for _, num := range *numbers {
		transformed = append(transformed, transform(num))
	}

	return transformed
}

// A function that returns a function based on the enum type passed as an argument
func transformerFunction(funcType functionType) numberTransform {
	switch funcType {
	case Double:
		return doubleNumber
	case Triple:
		return tripleNumber
	default:
		return nil
	}
}

func doubleNumber(number int) int {
	return number * 2
}

func tripleNumber(number int) int {
	return number * 3
}

// A function that returns a function using closures and creates a transformer function(a multiplication function) based on the value passed
func transformerCreator(value int) numberTransform {
	return func(number int) int {
		return number * value
	}
}

// Example of a recursive function to calculate factorial of a number
func recursiveFactorial(n int) int {
	if n <= 0 {
		return 1
	}

	return n * recursiveFactorial(n-1)
}

// function that uses variadic parameters
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
