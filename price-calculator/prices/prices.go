package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	for _, price := range job.InputPrices {
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * job.TaxRate
	}

	fmt.Printf("Tax Included Prices for %.2f \r\n", job.TaxRate)
	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       inputPrices,
		TaxIncludedPrices: make(map[string]float64),
	}
}
