package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// wordToken keeps source boundaries so callers can preserve punctuation without
// maintaining a second definition of what constitutes a word.
type wordToken struct {
	value string
	start int
	end   int
}

// clampRange prevents user-provided rune offsets from producing out-of-range slices.
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

// runeSubstring slices by rune offsets so multibyte text is never split mid-rune.
func runeSubstring(s string, start, end int) string {
	runes := []rune(s)
	start, end = clampRange(start, end, len(runes))
	if start >= end {
		return ""
	}
	return string(runes[start:end])
}

// runeIndex translates string-search byte offsets into rune offsets; empty or
// missing searches use -1 to honor the package-wide search contract.
func runeIndex(s, sub string, last bool) int {
	if sub == "" {
		return -1
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

// tokenizeWords applies the package's shared Unicode and acronym boundary rules.
// Keeping byte spans alongside values lets Words retain the exact source prefix.
func tokenizeWords(s string) []wordToken {
	runes := []rune(s)
	if len(runes) == 0 {
		return nil
	}

	byteOffsets := make([]int, 0, len(runes)+1)
	for byteOffset := range s {
		byteOffsets = append(byteOffsets, byteOffset)
	}
	byteOffsets = append(byteOffsets, len(s))

	var tokens []wordToken
	start := -1
	flush := func(end int) {
		if start < 0 || start == end {
			return
		}
		tokens = append(tokens, wordToken{
			value: s[byteOffsets[start]:byteOffsets[end]],
			start: byteOffsets[start],
			end:   byteOffsets[end],
		})
	}

	for i, r := range runes {
		isWord := unicode.IsLetter(r) || unicode.IsDigit(r)
		if !isWord && !(start >= 0 && unicode.IsMark(r)) {
			flush(i)
			start = -1
			continue
		}
		if start < 0 {
			start = i
			continue
		}
		if startsNewWord(runes, start, i) {
			flush(i)
			start = i
		}
	}
	flush(len(runes))

	return tokens
}

// startsNewWord recognizes lower/digit-to-upper transitions and the final
// capital of an acronym when it begins a conventionally cased word.
func startsNewWord(runes []rune, start, current int) bool {
	if current <= start || !isUpperWordRune(runes[current]) {
		return false
	}

	previous := current - 1
	for previous >= start && unicode.IsMark(runes[previous]) {
		previous--
	}
	if previous < start {
		return false
	}
	if unicode.IsLower(runes[previous]) || unicode.IsDigit(runes[previous]) {
		return true
	}
	if !isUpperWordRune(runes[previous]) {
		return false
	}

	for next := current + 1; next < len(runes); next++ {
		if unicode.IsMark(runes[next]) {
			continue
		}
		return unicode.IsLower(runes[next])
	}
	return false
}

// isUpperWordRune treats Unicode titlecase letters as capitals for word boundaries.
func isUpperWordRune(r rune) bool {
	return unicode.IsUpper(r) || unicode.IsTitle(r)
}

// wordTokenValues keeps casing transformations focused on text while the tokenizer
// retains source spans for boundary-sensitive operations.
func wordTokenValues(tokens []wordToken) []string {
	if tokens == nil {
		return nil
	}
	words := make([]string, len(tokens))
	for i, token := range tokens {
		words[i] = token.value
	}
	return words
}
