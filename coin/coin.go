package coin

import (
	"encoding/json"
	"strings"
	"log"
	"fmt"
	// "net"
	"github.com/gorilla/websocket"
)

func Price(ticker string) ([]byte, error) {
	// origin := fmt.Sprintf("/%s@ticker", ticker)

	url := fmt.Sprintf("wss://stream.binance.com:9443/ws/!ticker@%s", strings.ToLower(ticker))

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	print(ws)

	if err != nil {
		log.Fatal(err)
	}

	_, message, err := ws.ReadMessage()

	j, err := json.Marshal(message)

	print(j)

	return j, nil
}
