package str

import (
	"strings"
	"unicode"
)

// ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.
// @group Case
//
// Example: title map runes
//
//	v := str.Of("ß")
//	godump.Dump(v.ToTitle().String())
//	// #string SS
func (s String) ToTitle() String {
	var b strings.Builder
	for _, r := range s.s {
		if r == 'ß' {
			b.WriteString("SS")
			continue
		}
		b.WriteRune(unicode.ToTitle(r))
	}
	return String{s: b.String()}
}
