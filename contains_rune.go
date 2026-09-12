package str

import "strings"

// ContainsRune reports whether r occurs in the string.
// @group Search
//
// Example: ContainsRune
//
//	v := str.Of("café").ContainsRune('é')
//	println(v)
//	// #bool true
func (s String) ContainsRune(r rune) bool {
	return strings.ContainsRune(s.s, r)
}
