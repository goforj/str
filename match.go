package str

import "path"

// Match reports whether the entire string matches pattern using [path.Match] syntax.
// A malformed pattern returns an error, and wildcards do not match a slash.
// @group Match
//
// Example: match a shell pattern
//
//	matched, err := str.Of("billing:reports").Match("billing:*")
//	println(matched, err == nil)
//	// #bool true
//	// #bool true
func (s String) Match(pattern string) (bool, error) {
	return path.Match(pattern, s.s)
}
