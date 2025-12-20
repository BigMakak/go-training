package main

import (
	"errors"
	"fmt"
	"os"
)

const resultFile = "results.txt"

func main() {
	revenue := getUserInput("Enter your revenue amount:")
	expenses := getUserInput("Enter your expenses amount:")
	taxRate := getUserInput("Enter your tax rate percentage:")

	earningsBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Println("Earnings Before Tax: $", earningsBeforeTax)
	fmt.Println("Profit: $", profit)
	fmt.Println("Ration: $", ratio)

	values := []float64{earningsBeforeTax, profit, ratio}
	err := saveProfitValues(values)

	if err != nil {
		fmt.Println("Error saving profit values: ", err)
	} else {
		fmt.Println("Profit values saved successfully to", resultFile)
	}
}

func getUserInput(question string) float64 {
	var scanValue float64
	fmt.Print(question)
	_, err := fmt.Scan(&scanValue)

	if err != nil {
		fmt.Println("Error when getting user input: ", err)
		panic("Cannot process user input, terminating program...")
	}

	if scanValue <= 0 {
		fmt.Println("Input must be a positive number")
		panic("Input must be a positive number, terminating program...")
	}

	return scanValue
}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func saveProfitValues(values []float64) error {
	var resultText string = ""
	for _, value := range values {
		resultText += fmt.Sprintf("%f", value) + "\n"
	}

	err := writeResultToFile(resultText)
	if err != nil {
		return err
	}

	return nil
}

func writeResultToFile(value string) error {
	err := os.WriteFile(resultFile, []byte(value), 0644)
	if err != nil {
		return errors.New("error writing results to file")
	}

	return nil
}
