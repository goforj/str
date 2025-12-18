package str

import "unicode"

// UcFirst returns the string with the first rune upper-cased.
// @group Case
//
// Example: uppercase first rune
//
//	val := str.Of("gopher")
//	godump.Dump(val.UcFirst().String())
//
//	// #string Gopher
func (s String) UcFirst() String {
	runes := []rune(s.s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToUpper(runes[0])
	return String{s: string(runes)}
}
