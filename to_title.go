package str

import "strings"

// ToTitle maps every rune to Unicode titlecase.
// @group Case
//
// Example: ToTitle
//
//	v := str.Of("go").ToTitle().String()
//	println(v)
//	// #string GO
func (s String) ToTitle() String {
	return String{s: strings.ToTitle(s.s)}
}
