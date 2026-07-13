package str

import (
	"strings"
	"unicode"
)

// Slug returns a lowercase Unicode slug separated by hyphens.
// Unicode letters and digits are preserved, while every other run is collapsed
// to one hyphen.
// Similar: Kebab.
// @group Slug
//
// Example: build slug
//
//	v := str.Of("Go Forj Toolkit").Slug().String()
//	println(v)
//	// #string go-forj-toolkit
func (s String) Slug() String {
	var b strings.Builder
	b.Grow(len(s.s))
	pendingHyphen := false
	for _, r := range s.s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if pendingHyphen && b.Len() > 0 {
				b.WriteByte('-')
			}
			b.WriteRune(unicode.ToLower(r))
			pendingHyphen = false
			continue
		}
		if b.Len() > 0 {
			pendingHyphen = true
		}
	}

	return String{s: b.String()}
}
