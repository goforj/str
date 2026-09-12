package str

import "strings"

// CutLast splits around the last occurrence of sep.
// An empty sep is a match. On no match it returns the original string, "", false.
// CutLast preserves the Go 1.27 contract without raising the Go 1.24 minimum.
// @group Substrings
//
// Example: split a value
//
//	before, after, found := str.Of("go:forj").CutLast(":")
//	fmt.Println(before, after, found)
//	// #string go forj true
func (s String) CutLast(sep string) (before, after string, found bool) {
	index := strings.LastIndex(s.s, sep)
	if index < 0 {
		return s.s, "", false
	}
	return s.s[:index], s.s[index+len(sep):], true
}
