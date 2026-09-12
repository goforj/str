package str

import "strings"

// HasSuffix reports whether the string ends with suffix using a case-sensitive comparison.
// An empty suffix always matches.
// Similar: HasSuffixFold and HasPrefix.
// @group Search
//
// Example: has suffix
//
//	v := str.Of("gopher").HasSuffix("her")
//	println(v)
//	// #bool true
func (s String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.s, suffix)
}

// HasSuffixFold reports whether the string ends with suffix using Unicode simple case folding.
// An empty suffix always matches.
// Similar: HasSuffix and HasPrefixFold.
// @group Search
//
// Example: has suffix (case-insensitive)
//
//	v := str.Of("gopher").HasSuffixFold("HER")
//	println(v)
//	// #bool true
func (s String) HasSuffixFold(suffix string) bool {
	if suffix == "" {
		return true
	}
	start, ok := foldSuffixStart(s.s, suffix)
	if !ok {
		return false
	}

	end, ok := foldMatchAt(s.s, suffix, start)
	return ok && end == len(s.s)
}
