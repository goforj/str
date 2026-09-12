package str

import "strings"

// TrimFunc removes leading and trailing runes satisfying f.
// @group Cleanup
//
// Example: TrimFunc
//
//	v := str.Of("12Go34").TrimFunc(unicode.IsDigit).String()
//	println(v)
//	// #string Go
func (s String) TrimFunc(f func(rune) bool) String {
	return String{s: strings.TrimFunc(s.s, f)}
}
