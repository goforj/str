package str

import "strings"

// Kebab converts the string to kebab-case.
// @group Case
//
// Example: kebab case
//
//	val := str.Of("fooBar baz")
//	godump.Dump(val.Kebab().String())
//
//	// #string foo-bar-baz
func (s String) Kebab() String {
	words := splitWordsRunes(s.s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return String{s: strings.Join(words, "-")}
}
