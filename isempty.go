package str

import "unicode"

// IsEmpty reports whether the string has zero length.
// @group Checks
//
// Example: empty check
//
//	val := str.Of("")
//	godump.Dump(val.IsEmpty())
//
//	// #bool true
func (s String) IsEmpty() bool {
	return len(s.s) == 0
}

// IsBlank reports whether the string contains only Unicode whitespace.
// @group Checks
//
// Example: blank check
//
//	val := str.Of("  \\t\\n")
//	godump.Dump(val.IsBlank())
//
//	// #bool true
func (s String) IsBlank() bool {
	for _, r := range s.s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
