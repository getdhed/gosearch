package main

import (
	"1dz/gosearch/pkg/crawler"
	"1dz/gosearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"strings"
)

func main() {
	searchFlag := flag.String("s", "", "word to search")
	flag.Parse()
	searchWord := *searchFlag

	bot := spider.New()
	urls1, err := bot.Scan("https://go.dev/", 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("url from first:", urls1)
	urls2, err := bot.Scan("https://golang.org/", 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("url from second:", urls2)
	res := make([]crawler.Document, 0, len(urls1)+len(urls2))
	res = append(res, urls1...)
	res = append(res, urls2...)

	if searchWord == "" {
		return
	}

	for _, doc := range res {
		if strings.Contains(doc.URL, "s") {
			fmt.Println(doc.URL)
		}
	}
}
