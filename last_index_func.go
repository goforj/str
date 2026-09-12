package str

import "strings"

// LastIndexFunc returns the byte offset of the last rune satisfying f, or -1 if absent.
// @group Search
//
// Example: LastIndexFunc
//
//	v := str.Of("go2").LastIndexFunc(unicode.IsDigit)
//	println(v)
//	// #int 2
func (s String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(s.s, f)
}
