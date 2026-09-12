package str

import "strings"

// Clone copies the underlying bytes into a fresh allocation.
// Use it to release a large backing string retained by a small substring.
// Empty input returns an empty string without allocating.
// @group Transform
//
// Example: Clone
//
//	v := str.Of("gopher").Clone().String()
//	println(v)
//	// #string gopher
func (s String) Clone() String {
	return String{s: strings.Clone(s.s)}
}
