package main

import (
	"bot/irc"
	"bot/news"
	"bot/nft"
	// "bot/stocks"
	"bot/openai"
	"fmt"
	"os"
	"strings"
)

func handler(i *irc.IRC, m irc.Message) {
	//FIXME: if theres no input, bot crashes
	var q string = ""
	args := strings.Fields(m.RawArgs)
	arg := args[0]
	if len(args) > 1 {
		q = strings.Join(args[1:len(args)], " ")
	}

	if strings.Contains(arg, "!floor") {
		floor, err := nft.Search(q)
		if err != nil {
			fmt.Sprintf("%s", err)
			// panic(err)
		}

		if len(floor) > 1 {
			i.Say(m.Target, floor)
		}
	}
	if strings.Contains(arg, "!news") {
		// query := strings.ReplaceAll(q, " ", "")

		n, err := news.Search(q)
		if err != nil {
			fmt.Sprintf("%s", err)
		}

		if len(n.Articles) > 0 {
			for x := 0; x < len(n.Articles); x++ {
				i.Say(m.Target, fmt.Sprintf("%s - %s", n.Articles[x].Title, n.Articles[x].URL))
			}
		}
	}

	// if strings.Contains(arg, "!s") {
	// 	stock, err := stocks.Search(q)
	// 	if err != nil {
	// 		fmt.Sprint("%s", err)
	// 	}

	// 	if len(stock) > 1 {
	// 		i.Say(m.Target, stock)
	// 	}
	// }

	if strings.Contains(arg, "!tldr") {
		summary, err := openai.Summarize(q)
		if err != nil {
			fmt.Sprint("%s", err)
		}
		println(summary)

		if len(summary) > 1 {
			i.Say(m.Target, summary)
		}
	}
}
func main() {
	i := irc.New(&irc.Options{
		Addr: os.Getenv("SERVER"),
		Nick: os.Getenv("NICK"),
		Channels: []string{os.Getenv("CHANNELS")},
	})
	i.Register("PRIVMSG", handler)
	i.Start()
}
