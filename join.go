package str

import "strings"

// Join joins the provided words with sep.
// @group Words
//
// Example: join words
//
//	val := str.Of("unused")
//	godump.Dump(val.Join([]string{"foo", "bar"}, "-").String())
//
//	// #string foo-bar
func (s String) Join(words []string, sep string) String {
	return String{s: strings.Join(words, sep)}
}
