package str

import "strings"

// IndexAny returns the byte offset of the first rune in chars, or -1 if absent.
// @group Search
//
// Example: IndexAny
//
//	v := str.Of("go").IndexAny("o")
//	println(v)
//	// #int 1
func (s String) IndexAny(chars string) int {
	return strings.IndexAny(s.s, chars)
}
