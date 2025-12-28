package str

import "unicode"

// LcFirst returns the string with the first rune lower-cased.
// @group Case
//
// Example: lowercase first rune
//
//	v := str.Of("Gopher").LcFirst().String()
//	fmt.Println(v)
//	// #string gopher
func (s String) LcFirst() String {
	runes := []rune(s.s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToLower(runes[0])
	return String{s: string(runes)}
}
