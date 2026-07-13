package str

import "unicode"

// UcFirst returns the string with the first rune upper-cased.
// Similar: LcFirst and ToUpper.
// @group Case
//
// Example: uppercase first rune
//
//	v := str.Of("gopher").UcFirst().String()
//	println(v)
//	// #string Gopher
func (s String) UcFirst() String {
	runes := []rune(s.s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToUpper(runes[0])
	return String{s: string(runes)}
}
