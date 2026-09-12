package str

import "strings"

// ReplaceAll replaces all non-overlapping occurrences of old with new.
// An empty old matches at the beginning and after each UTF-8 sequence.
// @group Replace
//
// Example: ReplaceAll
//
//	v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
//	println(v)
//	// #string Go Gopher Go
func (s String) ReplaceAll(old, new string) String {
	return String{s: strings.ReplaceAll(s.s, old, new)}
}
