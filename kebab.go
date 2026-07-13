package str

import "strings"

// Kebab converts the string to kebab-case.
// Similar: Snake.
// @group Case
//
// Example: kebab case
//
//	v := str.Of("fooBar baz").Kebab().String()
//	println(v)
//	// #string foo-bar-baz
func (s String) Kebab() String {
	words := wordTokenValues(tokenizeWords(s.s))
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return String{s: strings.Join(words, "-")}
}
