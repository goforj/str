package str

import (
	"strings"
	"unicode"
)

// Initials returns the uppercase first rune of each detected word.
// Words are split the same way as SplitWords, including camel case and acronym boundaries.
// Similar: SplitWords.
// @group Words
//
// Example: collect word initials
//
//	v := str.Of("portableNetwork graphics").Initials().String()
//	println(v)
//	// #string PNG
func (s String) Initials() String {
	words := wordTokenValues(tokenizeWords(s.s))
	if len(words) == 0 {
		return String{s: ""}
	}

	var b strings.Builder
	b.Grow(len(words))
	for _, word := range words {
		for _, r := range word {
			b.WriteRune(unicode.ToUpper(r))
			break
		}
	}

	return String{s: b.String()}
}
