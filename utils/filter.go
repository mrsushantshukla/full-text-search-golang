package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowerCaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
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

func stopWordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	stopWords := map[string]struct{}{
		"a": {}, "an": {}, "the": {}, "is": {}, "in": {}, "at": {}, "of": {}, "on": {}, "and": {}, "or": {}}
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}
