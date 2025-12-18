package str

import (
	"strings"
	"unicode"
)

// Pascal converts the string to PascalCase.
// @group Case
//
// Example: pascal case
//
//	v := str.Of("foo_bar baz")
//	godump.Dump(v.Pascal().String())
//	// #string FooBarBaz
func (s String) Pascal() String {
	words := splitWordsRunes(s.s)
	for i, w := range words {
		if w == "" {
			continue
		}
		runes := []rune(strings.ToLower(w))
		runes[0] = unicode.ToTitle(runes[0])
		words[i] = string(runes)
	}
	return String{s: strings.Join(words, "")}
}
