package str

import "strings"

// ToLower returns a lowercase copy of the string using Unicode rules.
// @group Case
//
// Example: lowercase text
//
//	val := str.Of("GoLang")
//	godump.Dump(val.ToLower().String())
//
//	// #string golang
func (s String) ToLower() String {
	return String{s: strings.ToLower(s.s)}
}
