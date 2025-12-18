package str

import "strings"

// EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).
// @group Search
//
// Example: ends with any
//
//	val := str.Of("gopher")
//	godump.Dump(val.EndsWith("her", "cat"))
//
//	// #bool true
func (s String) EndsWith(suffixes ...string) bool {
	if len(suffixes) == 0 {
		return false
	}
	for _, suf := range suffixes {
		if strings.HasSuffix(s.s, suf) {
			return true
		}
	}
	return false
}

// EndsWithFold reports whether the string ends with any of the provided suffixes using Unicode case folding.
// @group Search
//
// Example: ends with (case-insensitive)
//
//	val := str.Of("gopher")
//	godump.Dump(val.EndsWithFold("HER"))
//
//	// #bool true
func (s String) EndsWithFold(suffixes ...string) bool {
	if len(suffixes) == 0 {
		return false
	}
	lower := strings.ToLower(s.s)
	for _, suf := range suffixes {
		if strings.HasSuffix(lower, strings.ToLower(suf)) {
			return true
		}
	}
	return false
}
