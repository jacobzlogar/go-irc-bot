package openai

import (
	"fmt"
	// "context"
	"os"
	"net/http"
	// "github.com/PullRequestInc/go-gpt3"
)

func Summarize(query string) (string, error) {
	key := os.Getenv("OPENAI_KEY")

	println(key)
	// ctx := context.Background()

	// client := gpt3.NewClient(key)

	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", query)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)

	println(res)

	return "", nil

	// if err != nil {
	// 	return "", err
	// }

	// resp, err := client.Completion(ctx, gpt3.CompletionRequest{
	// 	Prompt:    []string{summit},
	// 	MaxTokens: gpt3.IntPtr(30),
	// 	Stop:      []string{"."},
	// 	Echo:      true,
	// })

	// println(fmt.Sprintf("%", resp))

	// if err != nil {
	// 	println(err)
	// 	return "", err
	// }

	// return resp.Choices[0].Text, err
}
