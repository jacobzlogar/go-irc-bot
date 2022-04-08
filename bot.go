package main

import (
	"os"
	"fmt"
	"bot/irc"
	"bot/nft"
	"bot/news"
	// "bot/stocks"
	"strings"
)

func handler(i *irc.IRC, m irc.Message) {
	args := strings.Fields(m.RawArgs)
	arg := args[0]
	q := args[1]
	if strings.Contains(arg, "!floor") {
		if err := i.Say(m.Target, nft.Search(q)); err != nil {
			println(q)
		}
	}
	if strings.Contains(arg, "!news") {
		n := news.Search(q)
		for x := 0; x < 3; x++ {
			i.Say(m.Target, fmt.Sprintf("%s - %s", n.Articles[x].Title, n.Articles[x].URL))
		}
	}
	// if strings.Contains(s, "!s") {
	// 	if err := i.Say(m.Target, stocks.Search(args[x+1])); err != nil {
	// 		println(x)
	// 	}
	// }
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
