package str

import (
	"iter"
	"strings"
)

// FieldsFuncSeq returns an iterator over fields separated by runes satisfying f.
// Each iteration starts again from the beginning of the string.
// The predicate must return the same result for a given rune; its call order is unspecified.
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
