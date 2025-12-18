package str

import "strings"

// ToUpper returns an uppercase copy of the string using Unicode rules.
// @group Case
//
// Example: uppercase text
//
//	val := str.Of("GoLang")
//	godump.Dump(val.ToUpper().String())
//
//	// #string GOLANG
func (s String) ToUpper() String {
	return String{s: strings.ToUpper(s.s)}
}
