package str

import "strings"

// IndexByte returns the byte offset of the first byte equal to c, or -1 if absent.
// @group Search
//
// Example: IndexByte
//
//	v := str.Of("go").IndexByte('o')
//	println(v)
//	// #int 1
func (s String) IndexByte(c byte) int {
	return strings.IndexByte(s.s, c)
}
