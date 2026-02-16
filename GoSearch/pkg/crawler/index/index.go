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
	s = strings.TrimSpace(s)
	s = strings.Trim(s, " \t\r\n.,:;!?\"'()[]{}<>|\\/+-=*&#@%^`~")
	return s
}
func BuildRevIndexMap(res []crawler.Document) map[string][]int {
	idx := make(map[string][]int)

	for _, doc := range res {
		title := normalizeWord(doc.Title)
		words := strings.Fields(title)

		seen := make(map[string]struct{}, len(words))
		for _, w := range words {
			w = normalizeWord(w)
			if w == "" {
				continue
			}
			if _, ok := seen[w]; ok {
				continue
			}
			seen[w] = struct{}{}

			idx[w] = append(idx[w], doc.ID)
		}
	}

	return idx
}