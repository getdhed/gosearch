package main

import (
	"1dz/GoSearch/pkg/crawler"
	"1dz/GoSearch/pkg/crawler/index"
	fileWork "1dz/GoSearch/pkg/crawler/saveToFile"
	"1dz/GoSearch/pkg/crawler/spider"
	"fmt"
	"sort"
)

func main() {
	// searchFlag := flag.String("s", "", "word to search")
	// flag.Parse()
	// searchWord := *searchFlag
	// slc:=make([]index.Index,0)

	bot := spider.New()
	urls1, err := bot.Scan("https://go.dev/", 2)
	if err != nil {
		panic(err)
	}
	urls2, err := bot.Scan("https://golang.org/", 1)
	if err != nil {
		panic(err)
	}
	res := make([]crawler.Document, 0, len(urls1)+len(urls2))
	res = append(res, urls1...)
	res = append(res, urls2...)

	for i := range res {
		res[i].ID = i
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	rev := index.BuildRevIndex(res)
	for i := range rev {
		if len(rev[i].Value) > 0 {
			fmt.Println(rev[i])
		}

	}
	f:=fileWork.CreateFile()
	defer f.Close()
	fileWork.WriteDocuments(res,f)

}
