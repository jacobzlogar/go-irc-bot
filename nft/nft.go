package nft

import (
	"os"
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Opensea struct {
	Stats struct {
		OneDayVolume          float64 `json:"one_day_volume"`
		OneDayChange          float64 `json:"one_day_change"`
		OneDaySales           float64 `json:"one_day_sales"`
		OneDayAveragePrice    float64 `json:"one_day_average_price"`
		SevenDayVolume        float64 `json:"seven_day_volume"`
		SevenDayChange        float64 `json:"seven_day_change"`
		SevenDaySales         float64 `json:"seven_day_sales"`
		SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
		ThirtyDayVolume       float64 `json:"thirty_day_volume"`
		ThirtyDayChange       float64 `json:"thirty_day_change"`
		ThirtyDaySales        float64 `json:"thirty_day_sales"`
		ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
		TotalVolume           float64 `json:"total_volume"`
		TotalSales            float64 `json:"total_sales"`
		TotalSupply           float64 `json:"total_supply"`
		Count                 float64 `json:"count"`
		NumOwners             int     `json:"num_owners"`
		AveragePrice          float64 `json:"average_price"`
		NumReports            int     `json:"num_reports"`
		MarketCap             float64 `json:"market_cap"`
		FloorPrice            float64 `json:"floor_price"`
	} `json:"stats"`
}


type Nftstat struct {
	Hits []struct {
		Title       string  `json:"title"`
		Slug        string  `json:"slug"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Sales       float64 `json:"sales"`
		Formatted   struct {
			Title       string `json:"title"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Image       string `json:"image"`
			Sales       string `json:"sales"`
		} `json:"_formatted"`
	} `json:"hits"`
	NbHits           int    `json:"nbHits"`
	ExhaustiveNbHits bool   `json:"exhaustiveNbHits"`
	Query            string `json:"query"`
	Limit            int    `json:"limit"`
	Offset           int    `json:"offset"`
	ProcessingTimeMs int    `json:"processingTimeMs"`
}

func GetOpenSeaSlug(query string) string {
	request, err := json.Marshal(map[string]string{"q": query})

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://www.nft-stats.com/indexes/collections/search", bytes.NewBuffer(request))

	if err != nil {
		panic(err)
	}

	req.Header.Set("content-type", "application/json")

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", os.Getenv("NFTSTAT_KEY")))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	b := Nftstat{}

	json.Unmarshal([]byte(body), &b)

	return b.Hits[0].Slug
}

func Search(query string) string {
	q := GetOpenSeaSlug(query)

	url := fmt.Sprintf("https://api.opensea.io/api/v1/collection/%s/stats", q)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	b := Opensea{}

	json.Unmarshal([]byte(body), &b)

	floor := b.Stats.FloorPrice

	return fmt.Sprintf("floor of %s is at %s", query, strconv.FormatFloat(floor, 'f', 6, 64))
}
