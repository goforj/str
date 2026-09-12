package str

import "strings"

// Map maps each rune using mapping, dropping runes mapped to a negative value.
// @group Transform
//
// Example: Map
//
//	v := str.Of("go").Map(unicode.ToUpper).String()
//	println(v)
//	// #string GO
func (s String) Map(mapping func(rune) rune) String {
	return String{s: strings.Map(mapping, s.s)}
}
