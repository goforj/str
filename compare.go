package str

import "strings"

// Compare returns -1, 0, or 1 according to lexicographic byte order.
// @group Search
//
// Example: Compare
//
//	v := str.Of("go").Compare("rust")
//	println(v)
//	// #int -1
func (s String) Compare(other string) int {
	return strings.Compare(s.s, other)
}
