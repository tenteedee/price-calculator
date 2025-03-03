package cmdhandler

import (
	"fmt"
)

type CmdHandler struct {
}

func NewCmdHandler() *CmdHandler {
	return &CmdHandler{}
}

func (c *CmdHandler) ReadLines() ([]string, error) {
	fmt.Println("Enter your prices. Confirm with ENTER")
	var prices []string

	for {
		var price string
		fmt.Printf("Price %d: ", len(prices)+1)
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (c *CmdHandler) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
