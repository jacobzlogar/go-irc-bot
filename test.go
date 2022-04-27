package main

import (
	"bot/news"
	"bot/nft"
	"bot/stocks"
	"fmt"
)

func main() {
	test_stocks()
	test_news()
	test_nft()
}

func test_news() {
	res, err := news.Search("bitcoin")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	for i := 0; i < len(res.Articles); i++ {
		println(res.Articles[i].Title)
	}

	n, err := news.Search("")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	for i := 0; i < len(n.Articles); i++ {
		println(n.Articles[i].Title)
	}
}

func test_stocks() {
	res, err := stocks.Search("AAPL")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	println(res)
}

func test_nft() {
	res, err := nft.Search("moonbirds")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	println(res)
}
