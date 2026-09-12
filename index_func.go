package str

import "strings"

// IndexFunc returns the byte offset of the first rune satisfying f, or -1 if absent.
// @group Search
//
// Example: IndexFunc
//
//	v := str.Of("go2").IndexFunc(unicode.IsDigit)
//	println(v)
//	// #int 2
func (s String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(s.s, f)
}
