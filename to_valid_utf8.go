package str

import "strings"

// ToValidUTF8 replaces each run of invalid UTF-8 bytes with replacement.
// @group Transform
//
// Example: ToValidUTF8
//
//	v := str.Of("a\xff\xfeb").ToValidUTF8("?").String()
//	println(v)
//	// #string a?b
func (s String) ToValidUTF8(replacement string) String {
	return String{s: strings.ToValidUTF8(s.s, replacement)}
}
