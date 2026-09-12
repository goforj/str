package str

import "strings"

// Count counts non-overlapping occurrences of sub.
// An empty sub matches before and after each UTF-8 sequence.
// @group Search
//
// Example: Count
//
//	v := str.Of("gogophergo").Count("go")
//	println(v)
//	// #int 3
func (s String) Count(sub string) int {
	return strings.Count(s.s, sub)
}
