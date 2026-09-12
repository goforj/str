package str

import "strings"

// TrimRightFunc removes trailing runes satisfying f.
// @group Cleanup
//
// Example: TrimRightFunc
//
//	v := str.Of("12Go34").TrimRightFunc(unicode.IsDigit).String()
//	println(v)
//	// #string 12Go
func (s String) TrimRightFunc(f func(rune) bool) String {
	return String{s: strings.TrimRightFunc(s.s, f)}
}
