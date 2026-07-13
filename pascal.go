package str

import (
	"strings"
	"unicode"
)

// Pascal converts the string to PascalCase.
// Similar: Camel.
// @group Case
//
// Example: pascal case
//
//	v := str.Of("foo_bar baz").Pascal().String()
//	fmt.Println(v)
//	// #string FooBarBaz
func (s String) Pascal() String {
	words := wordTokenValues(tokenizeWords(s.s))
	for i, w := range words {
		runes := []rune(strings.ToLower(w))
		runes[0] = unicode.ToTitle(runes[0])
		words[i] = string(runes)
	}
	return String{s: strings.Join(words, "")}
}
