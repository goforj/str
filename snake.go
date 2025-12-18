package str

import "strings"

// Snake converts the string to snake_case using the provided separator (default "_").
// @group Case
//
// Example: snake case
//
//	val := str.Of("fooBar baz")
//	godump.Dump(val.Snake("_").String())
//
//	// #string foo_bar_baz
func (s String) Snake(sep string) String {
	if sep == "" {
		sep = "_"
	}
	words := splitWordsRunes(s.s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return String{s: strings.Join(words, sep)}
}
