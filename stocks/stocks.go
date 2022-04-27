package stocks

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Stock struct {
	Status     string  `json:"status"`
	From       string  `json:"from"`
	Symbol     string  `json:"symbol"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volume     int     `json:"volume"`
	AfterHours float64 `json:"afterHours"`
	PreMarket  float64 `json:"preMarket"`
}

func Search(query string) (string, error){
	key := os.Getenv("POLYGON_KEY")

	date := time.Now().AddDate(0, 0, -1)

	url := fmt.Sprintf("https://api.polygon.io/v1/open-close/%s/%s?adjusted=true&apiKey=%s",
		query,
		date.Format("2006-01-02"),
		key,
	)

	println(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	stock := Stock{}

	json.Unmarshal([]byte(body), &stock)

	if stock.Status == "NOT_FOUND" {
		return "Not found", nil
	}

	println(fmt.Sprint("%s", stock))

	open := stock.Open
	close := stock.Close
	after := stock.AfterHours

	msg := fmt.Sprintf("%s %s open: %v - close: %v - after hours: %v",
		query,
		date.Format("01-02-2006"),
		open,
		close,
		after,
	)

	return msg, nil
}
