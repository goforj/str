package str

import (
	"strings"
	"unicode"
)

// WrapWords wraps the string to the given rune width on whitespace boundaries,
// using breakStr between lines without discarding punctuation.
// Similar: Words.
// @group Words
//
// Example: wrap words
//
//	v := str.Of("The quick brown fox jumped over the lazy dog.").WrapWords(20, "\n").String()
//	println(v)
//	// #string The quick brown fox\njumped over the lazy\ndog.
func (s String) WrapWords(width int, breakStr string) String {
	if width <= 0 {
		return s
	}
	if breakStr == "" {
		breakStr = "\n"
	}

	runes := []rune(s.s)
	if len(runes) == 0 {
		return s
	}

	var out strings.Builder
	out.Grow(len(s.s))
	lineWidth := 0
	hasWord := false

	for pos := 0; pos < len(runes); {
		separatorStart := pos
		for pos < len(runes) && unicode.IsSpace(runes[pos]) {
			pos++
		}
		separator := runes[separatorStart:pos]
		if pos == len(runes) {
			out.WriteString(string(separator))
			break
		}

		wordStart := pos
		for pos < len(runes) && !unicode.IsSpace(runes[pos]) {
			pos++
		}
		word := runes[wordStart:pos]
		nextWidth, separatorBreaksLine := widthAfterWhitespace(lineWidth, separator)

		if hasWord && !separatorBreaksLine && nextWidth+len(word) > width {
			out.WriteString(breakStr)
			lineWidth = 0
		} else {
			out.WriteString(string(separator))
			lineWidth = nextWidth
		}
		out.WriteString(string(word))
		lineWidth += len(word)
		hasWord = true
	}

	return String{s: out.String()}
}

// widthAfterWhitespace preserves existing line boundaries while measuring all
// other whitespace as one rune each, matching the package's rune-safe width contract.
func widthAfterWhitespace(lineWidth int, whitespace []rune) (int, bool) {
	breaksLine := false
	for _, r := range whitespace {
		switch r {
		case '\n', '\r', '\u2028', '\u2029':
			lineWidth = 0
			breaksLine = true
		default:
			lineWidth++
		}
	}
	return lineWidth, breaksLine
}
