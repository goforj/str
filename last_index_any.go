package str

import "strings"

// LastIndexAny returns the byte offset of the last rune in chars, or -1 if absent.
// @group Search
//
// Example: LastIndexAny
//
//	v := str.Of("go").LastIndexAny("o")
//	println(v)
//	// #int 1
func (s String) LastIndexAny(chars string) int {
	return strings.LastIndexAny(s.s, chars)
}
