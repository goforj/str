package str

import (
	"iter"
	"strings"
)

// FieldsSeq returns a single-use iterator over fields separated by Unicode whitespace.
// Consecutive separators are combined; empty or separator-only input yields no fields.
// @group Split
//
// Example: FieldsSeq
//
//	v := slices.Collect(str.Of("a b c").FieldsSeq())
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) FieldsSeq() iter.Seq[string] {
	return strings.FieldsSeq(s.s)
}
