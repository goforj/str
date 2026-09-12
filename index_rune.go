package str

import "strings"

// IndexRune returns the byte offset of the first rune equal to r, or -1 if absent.
// @group Search
//
// Example: IndexRune
//
//	v := str.Of("go").IndexRune('o')
//	println(v)
//	// #int 1
func (s String) IndexRune(r rune) int {
	return strings.IndexRune(s.s, r)
}
