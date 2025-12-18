package str

import "strings"

// WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.
// @group Words
//
// Example: wrap words
//
//	v := str.Of("The quick brown fox jumped over the lazy dog.")
//	godump.Dump(v.WrapWords(20, "\n").String())
//	// #string The quick brown fox\njumped over the lazy\ndog.
func (s String) WrapWords(width int, breakStr string) String {
	if width <= 0 {
		return s
	}
	if breakStr == "" {
		breakStr = "\n"
	}

	words := splitWordsRunes(s.s)
	if len(words) == 0 {
		return s
	}

	var lines []string
	var current []string
	currentLen := 0

	flush := func() {
		lines = append(lines, strings.Join(current, " "))
		current = current[:0]
		currentLen = 0
	}

	for _, w := range words {
		if currentLen == 0 {
			current = append(current, w)
			currentLen = len([]rune(w))
			continue
		}
		if currentLen+1+len([]rune(w)) > width {
			flush()
			current = append(current, w)
			currentLen = len([]rune(w))
		} else {
			current = append(current, w)
			currentLen += 1 + len([]rune(w))
		}
	}
	flush()

	return String{s: strings.Join(lines, breakStr)}
}
