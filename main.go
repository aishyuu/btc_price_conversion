package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Exchanges struct {
	Btc struct {
		Usd float64 `json:"usd"`
		Jpy float64 `json:"jpy"`
		Gbp float64 `json:"gbp"`
		Eur float64 `json:"eur"`
		Cny float64 `json:"cny"`
	} `json:"btc"`
}

func main()  {
	res, err := http.Get("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/btc.json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Bitcoin price API is not working!")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var exchanges Exchanges

	err = json.Unmarshal(body, &exchanges)

	btc := exchanges.Btc

	fmt.Printf("USD - %.2f\n", btc.Usd)
	fmt.Printf("EUR - %.2f\n", btc.Eur)
	fmt.Printf("CNY - %.2f\n", btc.Cny)
	fmt.Printf("GBP - %.2f\n", btc.Gbp)
	fmt.Printf("JPY - %.2f\n", btc.Jpy)
}