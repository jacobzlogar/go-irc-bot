package news

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

func Search(query string) (News, error) {
	print(query)
	token := os.Getenv("NEWSAPI_KEY")

	to := time.Now()

	from := time.Now().AddDate(0, 0, -1)

	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&searchIn=title,description,content&from=%s&to=%s&pageSize=%s&apiKey=%s",
		query,
		from.Format("01-02-2006"),
		to.Format("01-02-2006"),
		strconv.Itoa(3),
		token,
	)

	news := News{}

	println(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return news, err
	}

	res, err := http.DefaultClient.Do(req)

	// println(res)

	if err != nil {
		return news, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &news)

	return news, nil
}
