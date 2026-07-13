package str

import "strings"

// Contains reports whether the string contains sub using a case-sensitive comparison.
// An empty substring is not a match.
// Similar: ContainsFold.
// @group Search
//
// Example: contains substring
//
//	v := str.Of("Go means gophers").Contains("gopher")
//	println(v)
//	// #bool true
func (s String) Contains(sub string) bool {
	return sub != "" && strings.Contains(s.s, sub)
}

// ContainsFold reports whether the string contains sub using Unicode simple case folding.
// An empty substring is not a match.
// Similar: Contains.
// @group Search
//
// Example: contains substring (case-insensitive)
//
//	v := str.Of("Go means gophers").ContainsFold("GOPHER")
//	println(v)
//	// #bool true
func (s String) ContainsFold(sub string) bool {
	_, _, ok := foldMatchRange(s.s, sub, 0)
	return ok
}
