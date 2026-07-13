package str

import "unicode"

// NormalizeSpace removes surrounding whitespace and collapses internal whitespace to single spaces.
// Similar: Trim.
// @group Cleanup
//
// Example: normalize whitespace
//
//	v := str.Of("  go   forj  ").NormalizeSpace().String()
//	println(v)
//	// #string go forj
func (s String) NormalizeSpace() String {
	var out []rune
	seenWord := false
	pendingSpace := false

	for _, r := range s.s {
		if unicode.IsSpace(r) {
			pendingSpace = seenWord
			continue
		}
		if pendingSpace {
			out = append(out, ' ')
			pendingSpace = false
		}
		out = append(out, r)
		seenWord = true
	}

	return String{s: string(out)}
}
