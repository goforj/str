package str

import "strings"

// ToUpper returns an uppercase copy of the string using Unicode rules.
// Similar: ToLower and UcFirst.
// @group Case
//
// Example: uppercase text
//
//	v := str.Of("GoLang").ToUpper().String()
//	println(v)
//	// #string GOLANG
func (s String) ToUpper() String {
	return String{s: strings.ToUpper(s.s)}
}
