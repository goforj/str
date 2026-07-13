package str

import "strings"

// Snake converts the string to snake_case.
// Similar: Kebab.
// @group Case
//
// Example: snake case
//
//	v := str.Of("fooBar baz").Snake().String()
//	println(v)
//	// #string foo_bar_baz
func (s String) Snake() String {
	words := wordTokenValues(tokenizeWords(s.s))
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return String{s: strings.Join(words, "_")}
}
