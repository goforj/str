package str

import "strings"

// Equals reports whether the string exactly matches other (case-sensitive).
// @group Comparison
//
// Example: exact match
//
//	v := str.Of("gopher").Equals("gopher")
//	godump.Dump(v)
//	// #bool true
func (s String) Equals(other string) bool {
	return s.s == other
}

// EqualsFold reports whether the string matches other using Unicode case folding.
// @group Comparison
//
// Example: case-insensitive match
//
//	v := str.Of("gopher").EqualsFold("GOPHER")
//	godump.Dump(v)
//	// #bool true
func (s String) EqualsFold(other string) bool {
	return strings.EqualFold(s.s, other)
}
