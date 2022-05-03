package stocks

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	// "os"
// 	"strings"
// 	// "time"
// )

// type Stock struct {
// 	QuoteResponse struct {
// 		Result []struct {
// 			Language                          string  `json:"language"`
// 			Region                            string  `json:"region"`
// 			QuoteType                         string  `json:"quoteType"`
// 			TypeDisp                          string  `json:"typeDisp"`
// 			QuoteSourceName                   string  `json:"quoteSourceName"`
// 			Triggerable                       bool    `json:"triggerable"`
// 			CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
// 			Currency                          string  `json:"currency"`
// 			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
// 			MarketCap                         int64   `json:"marketCap"`
// 			ForwardPE                         float64 `json:"forwardPE"`
// 			PriceToBook                       float64 `json:"priceToBook"`
// 			SourceInterval                    int     `json:"sourceInterval"`
// 			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
// 			PageViewGrowthWeekly              float64 `json:"pageViewGrowthWeekly"`
// 			AverageAnalystRating              string  `json:"averageAnalystRating"`
// 			Tradeable                         bool    `json:"tradeable"`
// 			MarketState                       string  `json:"marketState"`
// 			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
// 			PriceHint                         int     `json:"priceHint"`
// 			RegularMarketChange               float64 `json:"regularMarketChange"`
// 			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
// 			RegularMarketTime                 int     `json:"regularMarketTime"`
// 			RegularMarketPrice                float64 `json:"regularMarketPrice"`
// 			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
// 			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
// 			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
// 			RegularMarketVolume               int     `json:"regularMarketVolume"`
// 			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
// 			Bid                               float64 `json:"bid"`
// 			Ask                               float64 `json:"ask"`
// 			BidSize                           int     `json:"bidSize"`
// 			AskSize                           int     `json:"askSize"`
// 			FullExchangeName                  string  `json:"fullExchangeName"`
// 			FinancialCurrency                 string  `json:"financialCurrency"`
// 			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
// 			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
// 			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
// 			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
// 			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
// 			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
// 			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
// 			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
// 			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
// 			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
// 			DividendDate                      int     `json:"dividendDate"`
// 			EarningsTimestamp                 int     `json:"earningsTimestamp"`
// 			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
// 			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
// 			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
// 			TrailingPE                        float64 `json:"trailingPE"`
// 			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
// 			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
// 			EpsForward                        float64 `json:"epsForward"`
// 			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
// 			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
// 			SharesOutstanding                 int64   `json:"sharesOutstanding"`
// 			BookValue                         float64 `json:"bookValue"`
// 			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
// 			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
// 			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
// 			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
// 			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
// 			Exchange                          string  `json:"exchange"`
// 			ShortName                         string  `json:"shortName"`
// 			LongName                          string  `json:"longName"`
// 			MessageBoardID                    string  `json:"messageBoardId"`
// 			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
// 			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
// 			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
// 			Market                            string  `json:"market"`
// 			EsgPopulated                      bool    `json:"esgPopulated"`
// 			DisplayName                       string  `json:"displayName"`
// 			Symbol                            string  `json:"symbol"`
// 		} `json:"result"`
// 		Error interface{} `json:"error"`
// 	} `json:"quoteResponse"`
// }

// func Search(query string) (string, error){
// 	// key := os.Getenv("POLYGON_KEY")

// 	// date := time.Now().AddDate(0, 0, -1)

// 	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s",
// 		strings.ToUpper(query),
// 	)
// 	// url := fmt.Sprintf("https://api.polygon.io/v1/open-close/%s/%s?adjusted=true&apiKey=%s",
// 	// 	strings.ToUpper(query),
// 	// 	date.Format("2006-01-02"),
// 	// 	key,
// 	// )

// 	println(url)

// 	req, err := http.NewRequest("GET", url, nil)

// 	if err != nil {
// 		return "", err
// 	}

// 	req.Header.Add("Accept", "application/json")

// 	res, err := http.DefaultClient.Do(req)

// 	if err != nil {
// 		return "", err
// 	}

// 	defer res.Body.Close()

// 	body, _ := ioutil.ReadAll(res.Body)

// 	stock := Stock{}

// 	json.Unmarshal([]byte(body), &stock)

// 	// if stock.Status == "NOT_FOUND" {
// 	// 	return "Not found", nil
// 	// }

// 	println(fmt.Sprint("%s", stock))

// 	open := stock.Open
// 	close := stock.Close
// 	after := stock.AfterHours

// 	msg := fmt.Sprintf("%s %s open: %v - close: %v - after hours: %v",
// 		strings.ToUpper(query),
// 		open,
// 		close,
// 		after,
// 	)

// 	return msg, nil
// }
