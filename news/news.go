package news

import (
	"os"
	"fmt"
	"time"
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

func Search(query string) News {
	token := os.Getenv("NEWSAPI_KEY")

	to := time.Now()

	from := time.Now().AddDate(0, 0, -1)

	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&searchIn=title&from=%s&to=%s&apiKey=%s",
		query,
		from.Format("01-02-2006"),
		to.Format("01-02-2006"),
		token,
	)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	news := News{}

	json.Unmarshal([]byte(body), &news)

	return news
}
