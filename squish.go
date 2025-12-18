package str

import "unicode"

// Squish trims leading/trailing whitespace and collapses internal whitespace to single spaces.
// @group Cleanup
//
// Example: squish whitespace
//
//	v := str.Of("   go   forj  ")
//	godump.Dump(v.Squish().String())
//	// #string go forj
func (s String) Squish() String {
	var out []rune
	seenWord := false
	inSpace := false

	for _, r := range s.s {
		if unicode.IsSpace(r) {
			if seenWord && !inSpace {
				out = append(out, ' ')
			}
			inSpace = true
			continue
		}
		out = append(out, r)
		seenWord = true
		inSpace = false
	}

	// Trim trailing space inserted by collapse
	if len(out) > 0 && out[len(out)-1] == ' ' {
		out = out[:len(out)-1]
	}

	return String{s: string(out)}
}
