package str

import "strings"

// EqualFold reports whether the string matches other using Unicode simple case folding.
// @group Comparison
//
// Example: case-insensitive match
//
//	v := str.Of("gopher").EqualFold("GOPHER")
//	println(v)
//	// #bool true
func (s String) EqualFold(other string) bool {
	return strings.EqualFold(s.s, other)
}
