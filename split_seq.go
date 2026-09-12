package str

import (
	"iter"
	"strings"
)

// SplitSeq returns a single-use iterator over substrings separated by sep.
// An empty sep splits after each UTF-8 sequence.
// @group Split
//
// Example: SplitSeq
//
//	v := slices.Collect(str.Of("a,b,c").SplitSeq(","))
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) SplitSeq(sep string) iter.Seq[string] {
	return strings.SplitSeq(s.s, sep)
}
