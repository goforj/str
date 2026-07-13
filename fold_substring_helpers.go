package str

import (
	"strings"
	"unicode/utf8"
)

// foldMatchRange returns the first fold-equivalent, rune-aligned byte range at or after start.
// Rune-aligned candidates keep offsets valid when equivalent runes use different UTF-8 widths.
func foldMatchRange(s, sub string, start int) (int, int, bool) {
	if sub == "" || start < 0 || start >= len(s) {
		return 0, 0, false
	}

	for start < len(s) {
		end, ok := foldMatchAt(s, sub, start)
		if ok {
			return start, end, true
		}

		_, width := utf8.DecodeRuneInString(s[start:])
		start += width
	}

	return 0, 0, false
}

// foldMatchAt reports the byte end of a fold-equivalent candidate beginning at start.
// Counting sub's runes before EqualFold avoids assuming equivalent text has the same byte length.
func foldMatchAt(s, sub string, start int) (int, bool) {
	if sub == "" || start < 0 || start >= len(s) {
		return 0, false
	}

	end := start
	for range sub {
		if end >= len(s) {
			return 0, false
		}
		_, width := utf8.DecodeRuneInString(s[end:])
		end += width
	}

	return end, strings.EqualFold(s[start:end], sub)
}

// foldSuffixStart finds the only rune-aligned byte offset where suffix can match the end of s.
func foldSuffixStart(s, suffix string) (int, bool) {
	if suffix == "" {
		return 0, false
	}

	start := len(s)
	for range suffix {
		if start == 0 {
			return 0, false
		}
		_, width := utf8.DecodeLastRuneInString(s[:start])
		start -= width
	}

	return start, true
}
