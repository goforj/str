package str

import "strings"

// HasPrefix reports whether the string starts with prefix using a case-sensitive comparison.
// An empty prefix always matches.
// Similar: HasPrefixFold and HasSuffix.
// @group Search
//
// Example: has prefix
//
//	v := str.Of("gopher").HasPrefix("go")
//	println(v)
//	// #bool true
func (s String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.s, prefix)
}

// HasPrefixFold reports whether the string starts with prefix using Unicode simple case folding.
// An empty prefix always matches.
// Similar: HasPrefix and HasSuffixFold.
// @group Search
//
// Example: has prefix (case-insensitive)
//
//	v := str.Of("gopher").HasPrefixFold("GO")
//	println(v)
//	// #bool true
func (s String) HasPrefixFold(prefix string) bool {
	if prefix == "" {
		return true
	}
	_, ok := foldMatchAt(s.s, prefix, 0)
	return ok
}
