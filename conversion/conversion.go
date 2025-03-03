package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for index, value := range strings {
		floatPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {

			return nil, fmt.Errorf("invalid price at index %d", index)
		}

		floats[index] = floatPrice
	}

	return floats, nil
}
