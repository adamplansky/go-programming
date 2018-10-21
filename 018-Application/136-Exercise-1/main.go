package main

import (
	"encoding/json"
	"fmt"
)

type Currency struct {
	Name   string
	Amount float64
	Symbol string
}

func main() {
	currencies := []Currency{
		{"Bitcoin", 1.0001, "BTC"},
		{"Ethereum", 0.005, "ETH"},
	}
	b, err := json.Marshal(currencies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	testIsValid()

}
func testIsValid() {
	b := []byte(`[{
    "Name": "Bitcoin",
    "Amount": 1.0001,
    "Symbol": "BTC"
    }, {
      "Name": "Ethereum",
      "Amount": 0.005,
      "Symbol"
      "ETH"
      }]`)
	if json.Valid(b) == true {
		fmt.Printf("%v is valid json\n", string(b))
	} else {
		fmt.Printf("%v is invalid json\n", string(b))
	}
}
