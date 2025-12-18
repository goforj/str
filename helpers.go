package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func clampRange(start, end, length int) (int, int) {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}
	return start, end
}

func runeSubstring(s string, start, end int) string {
	runes := []rune(s)
	start, end = clampRange(start, end, len(runes))
	if start >= end {
		return ""
	}
	return string(runes[start:end])
}

func runeIndex(s, sub string, last bool) int {
	if sub == "" {
		return 0
	}
	var byteIdx int
	if last {
		byteIdx = strings.LastIndex(s, sub)
	} else {
		byteIdx = strings.Index(s, sub)
	}
	if byteIdx == -1 {
		return -1
	}
	return utf8.RuneCountInString(s[:byteIdx])
}

func splitWordsRunes(s string) []string {
	var words []string
	var buf []rune
	var prev rune

	flush := func() {
		if len(buf) == 0 {
			return
		}
		words = append(words, string(buf))
		buf = buf[:0]
	}

	for _, r := range s {
		isWord := unicode.IsLetter(r) || unicode.IsDigit(r)
		if !isWord {
			flush()
			prev = 0
			continue
		}

		// camelCase / PascalCase split: lower-to-upper transition starts new word.
		if len(buf) > 0 && unicode.IsUpper(r) && (unicode.IsLower(prev) || unicode.IsDigit(prev)) {
			flush()
		}

		buf = append(buf, r)
		prev = r
	}
	flush()
	return words
}
