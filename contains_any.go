package str

import "strings"

// ContainsAny reports whether any rune in chars occurs in the string.
// @group Search
//
// Example: ContainsAny
//
//	v := str.Of("gopher").ContainsAny("aeiou")
//	println(v)
//	// #bool true
func (s String) ContainsAny(chars string) bool {
	return strings.ContainsAny(s.s, chars)
}
