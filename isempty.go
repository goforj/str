package str

import "unicode"

// IsEmpty reports whether the string has zero length.
// Similar: IsBlank.
// @group Checks
//
// Example: empty check
//
//	v := str.Of("").IsEmpty()
//	println(v)
//	// #bool true
func (s String) IsEmpty() bool {
	return len(s.s) == 0
}

// IsBlank reports whether the string contains only Unicode whitespace.
// Similar: IsEmpty.
// @group Checks
//
// Example: blank check
//
//	v := str.Of("  \t\n")
//	println(v.IsBlank())
//	// #bool true
func (s String) IsBlank() bool {
	for _, r := range s.s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
