package main

import (
	// "bot/news"
	// "bot/nft"
	// "bot/stocks"
	// "bot/openai"
	"bot/wiki"
	"fmt"
)

func main() {
	test_wiki()
	// test_stocks()
	// test_news()
	// test_nft()
}

func test_wiki() {
	res, err := wiki.Search("bolsonaro")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	print(res)
	// println(fmt.Sprintf("%", res.Query.Search[0].Pageid))
}

// func test_openai() {
// 	res, err := openai.Summarize("dyson spheres")
// 	if err != nil {
// 		fmt.Printf("error %s", err)
// 	}
// 	print(res)
// }

// func test_news() {
// 	res, err := news.Search("the president of the united states of america")
// 	if err != nil {
// 		fmt.Printf("error %s", err)
// 	}
// 	for i := 0; i < len(res.Articles); i++ {
// 		println(res.Articles[i].Title)
// 	}
// }

// func test_stocks() {
// 	res, err := stocks.Search("AAPL")
// 	if err != nil {
// 		fmt.Printf("error %s", err)
// 	}
// 	println(res)
// }

// func test_nft() {
// 	res, err := nft.Search("moonbirds")
// 	if err != nil {
// 		fmt.Printf("error %s", err)
// 	}
// 	println(res)
// }
