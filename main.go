package main

import (
	"fmt"

	filehandler "github.com/tenteedee/price-calculator/file_handler"
	"github.com/tenteedee/price-calculator/prices"
)

func main() {
	inputPrices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		f := filehandler.NewFileHandler("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmd := cmdhandler.NewCmdHandler()
		job := prices.NewTaxIncludedPriceJob(f, taxRate, inputPrices)
		err := job.Process()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for taxRate, taxIncludedPrices := range results {
		fmt.Printf("%.2f ", taxRate)
		for _, taxIncludedPrice := range taxIncludedPrices {
			fmt.Printf("%.2f ", taxIncludedPrice)
		}
		fmt.Println()
	}
}
