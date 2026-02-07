package main

import (
	"1dz/GoSearch/pkg/crawler"
	"1dz/GoSearch/pkg/crawler/index"
	"1dz/GoSearch/pkg/crawler/spider"
	"fmt"
	"sort"
	"strings"
)

func normalizeWord(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, " \t\r\n.,:;!?\"'()[]{}<>|\\/+-=*&#@%^`~")
	return s
}
func BuildRevIndex(res []crawler.Document) []index.Index {
	var revIndex []index.Index

	for _, doc := range res {

		title := normalizeWord(doc.Title)
		words := strings.Fields(title)
		seen := make(map[string]struct{}, len(words))

		for _, w := range words {
			if w == "" {
				continue
			}
			if _, ok := seen[w]; ok {
				continue
			}
			seen[w] = struct{}{}
			found := false

			for i := range revIndex {
				if w == revIndex[i].Key {
					revIndex[i].Value = append(revIndex[i].Value, doc.ID)
					found = true
					break
				}
			}
			if !found {
				revIndex = append(revIndex, *index.NewIndex(w, doc.ID))
			}
		}
	}
	return revIndex
}
func main() {
	// searchFlag := flag.String("s", "", "word to search")
	// flag.Parse()
	// searchWord := *searchFlag

	// slc:=make([]index.Index,0)

	bot := spider.New()
	urls1, err := bot.Scan("https://go.dev/", 3)
	if err != nil {
		panic(err)
	}
	// for k, v := range urls1 {
	// 	slc[k].Key = v.Title
	// 	slc[k].Value =v.ID
	// }
	// fmt.Println("url from first:", urls1)
	urls2, err := bot.Scan("https://golang.org/", 1)
	if err != nil {
		panic(err)
	}
	// for k, v := range urls2 {
	// 	slc[k].Key = v.Title
	// 	slc[k].Value =v.ID
	// }

	// fmt.Println("url from second:", urls2)
	res := make([]crawler.Document, 0, len(urls1)+len(urls2))
	res = append(res, urls1...)
	res = append(res, urls2...)

	for i := range res {
		res[i].ID = i
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	rev := BuildRevIndex(res)
	for i := range rev {
		if len(rev[i].Value) > 0 {
			fmt.Println(rev[i])
		}

	}
	i := sort.Search(len(res), func(i int) bool {
		return res[i].ID >= id
	})

	if i < len(res) && res[i].ID == id {
		fmt.Println("found:", res[i])
	} else {
		fmt.Println("not found")
	}
	// if searchWord == "" {
	// 	return
	// }
	// for _, doc := range res {
	// 	if strings.Contains(doc.URL, searchWord) {
	// 		fmt.Println(doc.URL)
	// 	}
	// }
}
