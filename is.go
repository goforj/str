package str

import (
	"path"
	"strings"
)

// Is reports whether the string matches any of the provided wildcard patterns.
// Patterns use '*' as a wildcard. Case-sensitive.
// @group Match
//
// Example: wildcard match
//
//	v := str.Of("foo/bar").Is("foo/*")
//	godump.Dump(v)
//	// #bool true
func (s String) Is(patterns ...string) bool {
	if len(patterns) == 0 {
		return false
	}
	for _, p := range patterns {
		if p == "" {
			continue
		}
		if matchWildcard(s.s, p) {
			return true
		}
	}
	return false
}

func matchWildcard(s, pattern string) bool {
	ok, err := path.Match(pattern, s)
	if err != nil {
		// Fallback: try a simple contains if pattern is malformed.
		return strings.Contains(s, pattern)
	}
	return ok
}
