package main

import "testing"

func main() {
	news := news.Search("bitcoin")
	for i := 0; i <= 2; i++ {
		println(news.Articles[i].Title)
	}
	// println(news.Search("Bitcoin"))
}
