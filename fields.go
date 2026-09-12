package str

import "strings"

// Fields splits the string into fields separated by Unicode whitespace.
// Consecutive separators are combined; empty or separator-only input yields no fields.
// @group Split
//
// Example: Fields
//
//	v := str.Of("a b c").Fields()
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) Fields() []string {
	return strings.Fields(s.s)
}
