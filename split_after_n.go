package str

import "strings"

// SplitAfterN returns substrings including their trailing separator.
// A positive n limits the result to n substrings; zero returns nil and negative n has no limit.
// An empty sep splits after each UTF-8 sequence.
// @group Split
//
// Example: SplitAfterN
//
//	v := str.Of("a,b,c").SplitAfterN(",", 2)
//	fmt.Println(v)
//	// #[]string [a, b,c]
func (s String) SplitAfterN(sep string, n int) []string {
	return strings.SplitAfterN(s.s, sep, n)
}
