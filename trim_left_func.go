package str

import "strings"

// TrimLeftFunc removes leading runes satisfying f.
// @group Cleanup
//
// Example: TrimLeftFunc
//
//	v := str.Of("12Go34").TrimLeftFunc(unicode.IsDigit).String()
//	println(v)
//	// #string Go34
func (s String) TrimLeftFunc(f func(rune) bool) String {
	return String{s: strings.TrimLeftFunc(s.s, f)}
}
