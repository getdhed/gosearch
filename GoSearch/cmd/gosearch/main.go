package main

import (
	"1dz/GoSearch/pkg/crawler"
	"1dz/GoSearch/pkg/crawler/index"
	"1dz/GoSearch/pkg/crawler/spider"
	"1dz/netsrv"
	"log"
)

func main() {

	bot := spider.New()

	urls1, err := bot.Scan("https://go.dev/", 3)
	if err != nil {
		log.Fatal(err)
	}
	urls2, err := bot.Scan("https://golang.org/", 1)
	if err != nil {
		log.Fatal(err)
	}

	docs := make([]crawler.Document, 0, len(urls1)+len(urls2))
	docs = append(docs, urls1...)
	docs = append(docs, urls2...)

	for i := range docs {
		docs[i].ID = i
	}

	idx := index.BuildRevIndexMap(docs)

	docsByID := make(map[int]crawler.Document, len(docs))
	for _, d := range docs {
		docsByID[d.ID] = d
	}

	if err := netsrv.Serve(":8000", idx, docsByID); err != nil {
		log.Fatal(err)
	}
}
