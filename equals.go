package str

import "strings"

// Equals reports whether the string exactly matches other (case-sensitive).
// @group Comparison
//
// Example: exact match
//
//	val := str.Of("gopher")
//	godump.Dump(val.Equals("gopher"))
//
//	// #bool true
func (s String) Equals(other string) bool {
	return s.s == other
}

// EqualsFold reports whether the string matches other using Unicode case folding.
// @group Comparison
//
// Example: case-insensitive match
//
//	val := str.Of("gopher")
//	godump.Dump(val.EqualsFold("GOPHER"))
//
//	// #bool true
func (s String) EqualsFold(other string) bool {
	return strings.EqualFold(s.s, other)
}
