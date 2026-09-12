package str

import "strings"

// TrimRight removes trailing runes contained in cutset.
// @group Cleanup
//
// Example: TrimRight
//
//	v := str.Of("..GoForj!!").TrimRight(".!").String()
//	println(v)
//	// #string ..GoForj
func (s String) TrimRight(cutset string) String {
	return String{s: strings.TrimRight(s.s, cutset)}
}
