package str

import "unicode"

// IsAlpha reports whether the string contains at least one rune and every rune is a Unicode letter.
// @group Checks
//
// Example: alphabetic check
//
//	v := str.Of("Gopher").IsAlpha()
//	println(v)
//	// #bool true
func (s String) IsAlpha() bool {
	return allRunesMatch(s.s, unicode.IsLetter)
}
