package index

import (
	"1dz/GoSearch/pkg/crawler"
	"strings"
	"unicode"
)

func tokenize(s string) []string {
	s = strings.ToLower(s)

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}

	return strings.Fields(b.String())
}

func BuildRevIndexMap(res []crawler.Document) map[string][]int {
	idx := make(map[string][]int)

	for _, doc := range res {
		words := tokenize(doc.Title)

		seen := make(map[string]struct{}, len(words))

		for _, w := range words {
			if _, ok := seen[w]; ok {
				continue
			}
			seen[w] = struct{}{}

			idx[w] = append(idx[w], doc.ID)
		}
	}

	return idx
}
