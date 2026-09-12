package str

import (
	"iter"
	"strings"
)

// FieldsFuncSeq returns a single-use iterator over fields separated by runes satisfying f.
// Consecutive separators are combined; empty or separator-only input yields no fields.
// @group Split
//
// Example: FieldsFuncSeq
//
//	v := slices.Collect(str.Of("a b c").FieldsFuncSeq(unicode.IsSpace))
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) FieldsFuncSeq(f func(rune) bool) iter.Seq[string] {
	return strings.FieldsFuncSeq(s.s, f)
}
