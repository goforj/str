package str

import (
	"strings"
	"unicode"
)

// ToTitleSpecial maps every rune to Unicode titlecase using c's language-specific case rules.
// @group Case
//
// Example: ToTitleSpecial
//
//	v := str.Of("i").ToTitleSpecial(unicode.TurkishCase).String()
//	println(v)
//	// #string İ
func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return String{s: strings.ToTitleSpecial(c, s.s)}
}
