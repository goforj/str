package str

import (
	"strings"
	"unicode"
)

// Camel converts the string to camelCase.
// @group Case
//
// Example: camel case
//
//	val := str.Of("foo_bar baz")
//	godump.Dump(val.Camel().String())
//
//	// #string fooBarBaz
func (s String) Camel() String {
	words := splitWordsRunes(s.s)
	for i, w := range words {
		if w == "" {
			continue
		}
		if i == 0 {
			words[i] = strings.ToLower(w)
			continue
		}
		runes := []rune(strings.ToLower(w))
		runes[0] = unicode.ToTitle(runes[0])
		words[i] = string(runes)
	}
	return String{s: strings.Join(words, "")}
}
