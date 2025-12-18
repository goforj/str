package str

import "unicode"

// IsEmpty reports whether the string has zero length.
// @group Checks
//
// Example: empty check
//
//	v := str.Of("")
//	godump.Dump(v.IsEmpty())
//	// #bool true
func (s String) IsEmpty() bool {
	return len(s.s) == 0
}

// IsBlank reports whether the string contains only Unicode whitespace.
// @group Checks
//
// Example: blank check
//
//	v := str.Of("  \\t\\n")
//	godump.Dump(v.IsBlank())
//	// #bool true
func (s String) IsBlank() bool {
	for _, r := range s.s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
