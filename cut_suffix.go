package str

import "strings"

// CutSuffix removes suffix and reports whether it was present.
// An empty suffix is a match.
// @group Substrings
//
// Example: split a value
//
//	value, found := str.Of("go:forj").CutSuffix(":forj")
//	fmt.Println(value, found)
//	// #string go true
func (s String) CutSuffix(suffix string) (before string, found bool) {
	return strings.CutSuffix(s.s, suffix)
}
