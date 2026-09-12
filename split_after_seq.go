package str

import (
	"iter"
	"strings"
)

// SplitAfterSeq returns a single-use iterator over substrings including their trailing separator.
// An empty sep splits after each UTF-8 sequence.
// @group Split
//
// Example: SplitAfterSeq
//
//	v := slices.Collect(str.Of("a,b,c").SplitAfterSeq(","))
//	fmt.Println(v)
//	// #[]string [a, b, c]
func (s String) SplitAfterSeq(sep string) iter.Seq[string] {
	return strings.SplitAfterSeq(s.s, sep)
}
