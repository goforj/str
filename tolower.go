package str

import "strings"

// ToLower returns a lowercase copy of the string using Unicode rules.
// Similar: ToUpper and LcFirst.
// @group Case
//
// Example: lowercase text
//
//	v := str.Of("GoLang").ToLower().String()
//	println(v)
//	// #string golang
func (s String) ToLower() String {
	return String{s: strings.ToLower(s.s)}
}
