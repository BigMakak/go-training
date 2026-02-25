package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	values := []float64{10, 20, 30, 40}
	taxRates := []float64{1.0, 1.7, 1.1, 1.15}

	for _, taxRate := range taxRates {
		prices.NewTaxIncludedPriceJob(taxRate, values).Process()
	}

}
