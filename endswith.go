package str

import "strings"

// EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).
// @group Search
//
// Example: ends with any
//
//	v := str.Of("gopher").EndsWith("her", "cat")
//	godump.Dump(v)
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
//	v := str.Of("gopher").EndsWithFold("HER")
//	godump.Dump(v)
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
