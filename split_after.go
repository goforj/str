package str

import "strings"

// SplitAfter returns substrings including their trailing separator.
// An empty sep splits after each UTF-8 sequence.
// @group Split
//
// Example: SplitAfter
//
//	v := str.Of("a,b,c").SplitAfter(",")
//	fmt.Println(v)
//	// #[]string [a, b, c]
func (s String) SplitAfter(sep string) []string {
	return strings.SplitAfter(s.s, sep)
}
