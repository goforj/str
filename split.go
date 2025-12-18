package str

import "strings"

// Split splits the string by the given separator.
// @group Split
//
// Example: split on comma
//
//	v := str.Of("a,b,c").Split(",")
//	godump.Dump(v)
//	// #[]string [a b c]
func (s String) Split(sep string) []string {
	return strings.Split(s.s, sep)
}
