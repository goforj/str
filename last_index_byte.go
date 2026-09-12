package str

import "strings"

// LastIndexByte returns the byte offset of the last byte equal to c, or -1 if absent.
// @group Search
//
// Example: LastIndexByte
//
//	v := str.Of("go").LastIndexByte('o')
//	println(v)
//	// #int 1
func (s String) LastIndexByte(c byte) int {
	return strings.LastIndexByte(s.s, c)
}
