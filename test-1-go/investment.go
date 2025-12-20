package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investment := getUserInput("Enter the amount you want to invest:")

	expectedReturnRate := getUserInput("Enter the expected Return rate:")

	years := getUserInput("Enter the number of years you plan to invest:")

	futureValue, futureInflationAdjusted := calculateInvestment(investment, expectedReturnRate, years)

	investmentAmmount := fmt.Sprintf("Investment amount: %.0f$ \n", investment)
	fmt.Print(investmentAmmount)

	printInvesment(futureValue, futureInflationAdjusted)
	fmt.Println(`Multiple-line
			string literal example`)
}

func getUserInput(question string) (value float64) {
	fmt.Print(question)
	fmt.Scan(&value)
	return value
}

func printInvesment(futureValue float64, futureInflationAdjusted float64) {
	fmt.Printf("Future value of the investment: %.1f \n", futureValue)
	fmt.Printf("Future value adjusted for inflation: $%.2f \n", futureInflationAdjusted)
}

func calculateInvestment(investment float64, expectedReturnRate float64, years float64) (float64, float64) {
	futureValue := investment * math.Pow(1+expectedReturnRate/100, years)
	futureInflationAdjusted := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureInflationAdjusted
}
