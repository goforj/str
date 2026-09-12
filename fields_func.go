package str

import "strings"

// FieldsFunc splits the string into fields separated by runes satisfying f.
// Consecutive separators are combined; empty or separator-only input yields no fields.
// @group Split
//
// Example: FieldsFunc
//
//	v := str.Of("a b c").FieldsFunc(unicode.IsSpace)
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) FieldsFunc(f func(rune) bool) []string {
	return strings.FieldsFunc(s.s, f)
}
