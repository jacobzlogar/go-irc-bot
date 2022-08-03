package game

import (
	"bot/coin"
	"fmt"
)

func Create(ticker string, timeframe string) (string, error) {
	price, err := coin.Price(ticker)

	print(price)
	if err != nil {
		fmt.Sprintf("%s", err)
	}

	return price, nil
}
