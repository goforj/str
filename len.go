package str

import "unicode/utf8"

// RuneCount returns the number of Unicode code points in the string.
// @group Length
//
// Example: count runes instead of bytes
//
//	v := str.Of("gophers 🦫").RuneCount()
//	println(v)
//	// #int 9
func (s String) RuneCount() int {
	return utf8.RuneCountInString(s.s)
}
