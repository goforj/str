package str

import "strings"

// StartsWith reports whether the string starts with any of the provided prefixes (case-sensitive).
// @group Search
//
// Example: starts with any
//
//	val := str.Of("gopher")
//	godump.Dump(val.StartsWith("go", "rust"))
//
//	// #bool true
func (s String) StartsWith(prefixes ...string) bool {
	if len(prefixes) == 0 {
		return false
	}
	for _, p := range prefixes {
		if strings.HasPrefix(s.s, p) {
			return true
		}
	}
	return false
}

// StartsWithFold reports whether the string starts with any of the provided prefixes using Unicode case folding.
// @group Search
//
// Example: starts with (case-insensitive)
//
//	val := str.Of("gopher")
//	godump.Dump(val.StartsWithFold("GO"))
//
//	// #bool true
func (s String) StartsWithFold(prefixes ...string) bool {
	if len(prefixes) == 0 {
		return false
	}
	lower := strings.ToLower(s.s)
	for _, p := range prefixes {
		if strings.HasPrefix(lower, strings.ToLower(p)) {
			return true
		}
	}
	return false
}
