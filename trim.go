package str

import "strings"

// Trim removes leading and trailing runes contained in cutset.
// @group Cleanup
//
// Example: Trim
//
//	v := str.Of("..GoForj!!").Trim(".!").String()
//	println(v)
//	// #string GoForj
func (s String) Trim(cutset string) String {
	return String{s: strings.Trim(s.s, cutset)}
}
