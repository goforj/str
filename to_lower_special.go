package str

import (
	"strings"
	"unicode"
)

// ToLowerSpecial maps every rune to lowercase using c's language-specific case rules.
// @group Case
//
// Example: ToLowerSpecial
//
//	v := str.Of("I").ToLowerSpecial(unicode.TurkishCase).String()
//	println(v)
//	// #string ı
func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return String{s: strings.ToLowerSpecial(c, s.s)}
}
