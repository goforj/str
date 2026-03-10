package str

import "unicode"

// IsAlnum reports whether the string contains at least one rune and every rune is a Unicode letter or number.
// @group Checks
//
// Example: alphanumeric check
//
//	v := str.Of("Gopher2025").IsAlnum()
//	println(v)
//	// #bool true
func (s String) IsAlnum() bool {
	return allRunesMatch(s.s, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
}
