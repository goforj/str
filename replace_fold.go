package str

import "strings"

// ReplaceFold replaces all non-overlapping occurrences of old with repl using Unicode simple case folding.
// An empty old string leaves the receiver unchanged.
// Similar: ReplaceAll.
// @group Replace
//
// Example: replace all (case-insensitive)
//
//	v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
//	println(v)
//	// #string Go Gopher Go
func (s String) ReplaceFold(old, repl string) String {
	if old == "" {
		return s
	}
	out, ok := replaceFoldAll(s.s, old, repl)
	if !ok {
		return s
	}
	return String{s: out}
}

// replaceFoldAll builds a replacement only after the first match so missing values preserve the receiver.
func replaceFoldAll(s, old, repl string) (string, bool) {
	var b strings.Builder
	matched := false
	lastByte := 0
	searchByte := 0

	for {
		start, end, ok := foldMatchRange(s, old, searchByte)
		if !ok {
			break
		}

		if !matched {
			b.Grow(len(s))
		}
		b.WriteString(s[lastByte:start])
		b.WriteString(repl)
		lastByte = end
		searchByte = end
		matched = true
	}

	if !matched {
		return s, false
	}

	b.WriteString(s[lastByte:])
	return b.String(), true
}
