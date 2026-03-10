package str

import "unicode"

// IsNumeric reports whether the string contains at least one rune and every rune is a Unicode number.
// @group Checks
//
// Example: numeric check
//
//	v := str.Of("12345").IsNumeric()
//	println(v)
//	// #bool true
func (s String) IsNumeric() bool {
	return allRunesMatch(s.s, unicode.IsNumber)
}
