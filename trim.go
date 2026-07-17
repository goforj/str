package str

import (
	"strings"
	"unicode"
)

// Trim removes leading and trailing Unicode whitespace.
// Similar: TrimLeft, TrimRight, and TrimChars.
// @group Cleanup
//
// Example: trim whitespace
//
//	v := str.Of("  GoForj  ").Trim().String()
//	println(v)
//	// #string GoForj
func (s String) Trim() String {
	return String{s: strings.TrimSpace(s.s)}
}

// TrimChars removes leading and trailing runes contained in cutset.
// Similar: Trim.
// @group Cleanup
//
// Example: trim selected characters
//
//	v := str.Of("..GoForj!!").TrimChars(".!").String()
//	println(v)
//	// #string GoForj
func (s String) TrimChars(cutset string) String {
	return String{s: strings.Trim(s.s, cutset)}
}

// TrimLeft removes leading Unicode whitespace.
// Similar: Trim and TrimRight.
// @group Cleanup
//
// Example: trim left
//
//	v := str.Of("  GoForj  ").TrimLeft().String()
//	println(v)
//	// #string GoForj\u0020\u0020
func (s String) TrimLeft() String {
	return String{s: strings.TrimLeftFunc(s.s, unicode.IsSpace)}
}

// TrimRight removes trailing Unicode whitespace.
// Similar: Trim and TrimLeft.
// @group Cleanup
//
// Example: trim right
//
//	v := str.Of("  GoForj  ").TrimRight().String()
//	println(v)
//	// #string \u0020\u0020GoForj
func (s String) TrimRight() String {
	return String{s: strings.TrimRightFunc(s.s, unicode.IsSpace)}
}
