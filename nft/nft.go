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

func GetOpenSeaSlug(query string) (string, error) {
	request, err := json.Marshal(map[string]string{"q": query})

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://www.nft-stats.com/indexes/collections/search", bytes.NewBuffer(request))

	if err != nil {
		return "", err
	}
	println(os.Getenv("NFTSTAT_KEY"))

	req.Header.Set("content-type", "application/json")

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", os.Getenv("NFTSTATS_KEY")))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	// println(fmt.Sprintf("%s", body))

	b := Nftstat{}

	json.Unmarshal([]byte(body), &b)

	if len(b.Hits) > 1 {
		hit := b.Hits[0]
		return hit.Slug, err
	}

	return "", err
}

func Search(query string) (string, error) {
	slug, err := GetOpenSeaSlug(query)
	// println(slug)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.opensea.io/api/v1/collection/%s/stats", slug)

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

	b := Opensea{}

	json.Unmarshal([]byte(body), &b)

	floor := b.Stats.FloorPrice

	collection_url := fmt.Sprintf("https://opensea.io/collection/%s", slug)

	msg := fmt.Sprintf("%s floor %s ⟠: %s",
		query,
		strconv.FormatFloat(floor, 'f', 2, 64),
		collection_url,
	)

	return msg, nil
}
