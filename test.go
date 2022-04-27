package main

import (
	"bot/news"
	"bot/nft"
	"fmt"
)

func main() {
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


func test_nft() {
	res, err := nft.Search("moonbirds")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	println(res)
}
