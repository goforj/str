package str

import "unicode"

// UcSplit splits the string on uppercase boundaries and delimiters, returning segments.
// @group Split
//
// Example: split on upper-case boundaries
//
//	v := str.Of("HTTPRequestID").UcSplit()
//	godump.Dump(v)
//	// #[]string [HTTP Request ID]
func (s String) UcSplit() []string {
	runes := []rune(s.s)
	if len(runes) == 0 {
		return nil
	}

	var parts []string
	start := 0

	flush := func(pos int) {
		if pos-start > 0 {
			parts = append(parts, string(runes[start:pos]))
		}
	}

	for i := 1; i < len(runes); i++ {
		prev := runes[i-1]
		cur := runes[i]

		// Split on delimiters
		if !(unicode.IsLetter(cur) || unicode.IsDigit(cur)) {
			flush(i)
			start = i + 1
			continue
		}

		// Transition from lower/digit to upper starts new segment
		if unicode.IsUpper(cur) && (unicode.IsLower(prev) || unicode.IsDigit(prev)) {
			flush(i)
			start = i
			continue
		}

		// Acronym followed by lower: split before the lower
		if unicode.IsLower(cur) && unicode.IsUpper(prev) && i-start > 1 {
			flush(i - 1)
			start = i - 1
		}
	}
	flush(len(runes))

	return parts
}
