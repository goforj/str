package str

import "strings"

// Replace replaces the first n non-overlapping occurrences of old with new.
// A negative n replaces all matches; zero leaves the string unchanged.
// An empty old matches at the beginning and after each UTF-8 sequence.
// @group Replace
//
// Example: Replace
//
//	v := str.Of("go go go").Replace("go", "Go", 2).String()
//	println(v)
//	// #string Go Go go
func (s String) Replace(old, new string, n int) String {
	return String{s: strings.Replace(s.s, old, new, n)}
}
