package index

import (
	"1dz/GoSearch/pkg/crawler"
	"strings"
)

type Index struct {
	Key   string
	Value []int
}

func NewIndex(s string, id int) *Index {

	return &Index{
		Key:   s,
		Value: []int{},
	}
}
func normalizeWord(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, " \t\r\n.,:;!?\"'()[]{}<>|\\/+-=*&#@%^`~")
	return s
}
func BuildRevIndex(res []crawler.Document) []Index {
	var revIndex []Index

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
				revIndex = append(revIndex, *NewIndex(w, doc.ID))
			}
		}
	}
	return revIndex
}