package str

import "strings"

// Cut splits around the first occurrence of sep.
// An empty sep is a match. On no match it returns the original string, "", false.
// @group Substrings
//
// Example: split a value
//
//	before, after, found := str.Of("go:forj").Cut(":")
//	fmt.Println(before, after, found)
//	// #string go forj true
func (s String) Cut(sep string) (before, after string, found bool) {
	return strings.Cut(s.s, sep)
}
