package str

import "strings"

// CutPrefix removes prefix and reports whether it was present.
// An empty prefix is a match.
// @group Substrings
//
// Example: split a value
//
//	value, found := str.Of("go:forj").CutPrefix("go:")
//	fmt.Println(value, found)
//	// #string forj true
func (s String) CutPrefix(prefix string) (after string, found bool) {
	return strings.CutPrefix(s.s, prefix)
}
