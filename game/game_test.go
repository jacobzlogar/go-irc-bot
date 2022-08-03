package game_test

import (
	"bot/game"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	res, err := game.Create("btc", "1h")

	if err != nil {
		fmt.Sprintf("%s", err)
	}

	t.Logf("%s", res)
	// fmt.Sprintf("%s", res)
}
