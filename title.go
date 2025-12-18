package str

import (
	"strings"
	"unicode"
)

// Title converts the string to title case (first letter of each word upper, rest lower) using Unicode rules.
// @group Case
//
// Example: title case words
//
//	v := str.Of("a nice title uses the correct case")
//	godump.Dump(v.Title().String())
//	// #string A Nice Title Uses The Correct Case
func (s String) Title() String {
	var b strings.Builder
	b.Grow(len(s.s))

	prevIsWord := false
	for _, r := range s.s {
		isWord := unicode.IsLetter(r) || unicode.IsDigit(r)
		if isWord && !prevIsWord {
			b.WriteRune(unicode.ToTitle(r))
		} else if isWord {
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
		prevIsWord = isWord
	}

	return String{s: b.String()}
}
