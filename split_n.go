package str

import "strings"

// SplitN returns substrings separated by sep.
// A positive n limits the result to n substrings; zero returns nil and negative n has no limit.
// An empty sep splits after each UTF-8 sequence.
// @group Split
//
// Example: SplitN
//
//	v := str.Of("a,b,c").SplitN(",", 2)
//	fmt.Println(v)
//	// #[]string [a b,c]
func (s String) SplitN(sep string, n int) []string {
	return strings.SplitN(s.s, sep, n)
}
