package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}
func stopworpFilter(tokens []string) []string {
	var stopworps = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {}}
	// investigate make()
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopworps[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
