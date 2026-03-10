package str

import "strings"

// HasSurrounding reports whether the string starts with before and ends with after.
// If after is empty, before is used for both sides.
// @group Affixes
//
// Example: has matching delimiters
//
//	v := str.Of(`"GoForj"`).HasSurrounding(`"`, "")
//	println(v)
//	// #bool true
func (s String) HasSurrounding(before, after string) bool {
	if after == "" {
		after = before
	}

	if len(before)+len(after) > len(s.s) {
		return false
	}

	return strings.HasPrefix(s.s, before) && strings.HasSuffix(s.s, after)
}
