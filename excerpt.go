package str

import "strings"

// Excerpt returns a snippet around the first occurrence of needle with the given radius.
// If needle is not found, an empty string is returned. If radius <= 0, a default of 100 is used.
// Omission is used at the start/end when text is trimmed (default "...").
// @group Snippet
//
// Example: excerpt with radius
//
//	v := str.Of("This is my name").Excerpt("my", 3, "...")
//	godump.Dump(v.String())
//	// #string ...is my na...
func (s String) Excerpt(needle string, radius int, omission string) String {
	if needle == "" {
		return String{s: ""}
	}
	if radius <= 0 {
		radius = 100
	}
	if omission == "" {
		omission = "..."
	}

	idx := strings.Index(s.s, needle)
	if idx == -1 {
		return String{s: ""}
	}

	runes := []rune(s.s)
	needleRunes := []rune(needle)

	startRune := len([]rune(s.s[:idx]))
	endRune := startRune + len(needleRunes)

	lo := startRune - radius
	if lo < 0 {
		lo = 0
	}
	hi := endRune + radius
	if hi > len(runes) {
		hi = len(runes)
	}

	var out strings.Builder
	if lo > 0 {
		out.WriteString(omission)
	}
	out.WriteString(string(runes[lo:hi]))
	if hi < len(runes) {
		out.WriteString(omission)
	}

	return String{s: out.String()}
}
