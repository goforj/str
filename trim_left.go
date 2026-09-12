package str

import "strings"

// TrimLeft removes leading runes contained in cutset.
// @group Cleanup
//
// Example: TrimLeft
//
//	v := str.Of("..GoForj!!").TrimLeft(".!").String()
//	println(v)
//	// #string GoForj!!
func (s String) TrimLeft(cutset string) String {
	return String{s: strings.TrimLeft(s.s, cutset)}
}
