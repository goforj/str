package str

import "strings"

// ContainsFunc reports whether any rune satisfies f.
// @group Search
//
// Example: ContainsFunc
//
//	v := str.Of("go2").ContainsFunc(unicode.IsDigit)
//	println(v)
//	// #bool true
func (s String) ContainsFunc(f func(rune) bool) bool {
	return strings.ContainsFunc(s.s, f)
}
