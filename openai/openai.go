package openai

import (
	"os"
	// "fmt"
	// "strings"
	"context"
	"bot/wiki"
	// "net/http"
	// "io/ioutil"
	// "encoding/json"

	"github.com/PullRequestInc/go-gpt3"
)

func Summarize(query string) (string, error) {
	key := os.Getenv("OPENAI_KEY")

	println(key)

	extract, err := wiki.Search(query)

	if err != nil {
		return "", err
	}
	// query = strings.Title(query)

	// search := strings.ReplaceAll(query, " ", "_")

	print(extract)

	ctx := context.Background()

	client := gpt3.NewClient(key)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{extract},
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

	return "", nil

}
