package prices

import (
	"fmt"

	"github.com/tenteedee/price-calculator/conversion"
	iomanager "github.com/tenteedee/price-calculator/io_manager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:         io,
		TaxRate:           taxRate,
		InputPrices:       inputPrices,
		TaxIncludedPrices: make(map[string]string),
	}
}

func (t *TaxIncludedPriceJob) Process() error {
	err := t.LoadData()

	if err != nil {
		return err
	}

	for _, price := range t.InputPrices {
		priceStr := fmt.Sprintf("%.2f", price)
		t.TaxIncludedPrices[priceStr] = fmt.Sprintf("%.2f", (price * (1 + t.TaxRate)))
	}

	fmt.Println(t.TaxIncludedPrices)

	return t.IOManager.WriteResult(t)
}

func (t *TaxIncludedPriceJob) LoadData() error {
	data, err := t.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(data)

	if err != nil {
		return err
	}

	t.InputPrices = prices
	return nil
}
