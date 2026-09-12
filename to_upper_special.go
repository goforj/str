package str

import (
	"strings"
	"unicode"
)

// ToUpperSpecial maps every rune to uppercase using c's language-specific case rules.
// @group Case
//
// Example: ToUpperSpecial
//
//	v := str.Of("i").ToUpperSpecial(unicode.TurkishCase).String()
//	println(v)
//	// #string İ
func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return String{s: strings.ToUpperSpecial(c, s.s)}
}
