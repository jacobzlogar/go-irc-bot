package openai

import (
	"os"
	"fmt"
	"strings"
	"context"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/PullRequestInc/go-gpt3"
)

type Wiki struct {
	Batchcomplete bool `json:"batchcomplete"`
	Query         struct {
		Pages []struct {
			Pageid  int    `json:"pageid"`
			Ns      int    `json:"ns"`
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

func Summarize(query string) (string, error) {
	key := os.Getenv("OPENAI_KEY")

	println(key)
	query = strings.Title(query)

	search := strings.ReplaceAll(query, " ", "_")

	print(search)

	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&titles=%s&formatversion=2&exsentences=1&exlimit=1&explaintext=1",
		search,
	)

	print(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	wiki := Wiki{}

	json.Unmarshal([]byte(body), &wiki)

	ctx := context.Background()

	client := gpt3.NewClient(key)

	if len(wiki.Query.Pages) > 0 {
		page := wiki.Query.Pages[0].Extract
		resp, err := client.Completion(ctx, gpt3.CompletionRequest{
			Prompt:    []string{page},
			MaxTokens: gpt3.IntPtr(100),
			Stop:      []string{"."},
			Echo:      true,
		})

		if err != nil {
			return "", err
		}

		if len(resp.Choices) > 0 {
			return resp.Choices[0].Text, nil
		}
	}

	return "", nil

}
