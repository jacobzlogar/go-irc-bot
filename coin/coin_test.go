package coin_test

import (
	"bot/coin"
	"testing"
)

func TestPrice(t *testing.T) {
	price, err := coin.Price("btcusdt")

    if price != nil || err == nil {
        t.Fatalf(`Price("") = %q, %v, want "", error`, price, err)
    }

	print(price)
}
