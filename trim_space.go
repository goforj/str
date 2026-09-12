package str

import "strings"

// TrimSpace removes leading and trailing Unicode whitespace.
// @group Cleanup
//
// Example: TrimSpace
//
//	v := str.Of("  GoForj  ").TrimSpace().String()
//	println(v)
//	// #string GoForj
func (s String) TrimSpace() String {
	return String{s: strings.TrimSpace(s.s)}
}
