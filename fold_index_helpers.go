package str

import "unicode"

func runeIndexFold(s, sub string, last bool) int {
	if sub == "" {
		return 0
	}

	haystack := []rune(s)
	needle := []rune(sub)
	if len(needle) > len(haystack) {
		return -1
	}

	matchesAt := func(start int) bool {
		for i, r := range needle {
			if unicode.ToLower(haystack[start+i]) != unicode.ToLower(r) {
				return false
			}
		}
		return true
	}

	if last {
		for i := len(haystack) - len(needle); i >= 0; i-- {
			if matchesAt(i) {
				return i
			}
		}
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if matchesAt(i) {
			return i
		}
	}

	return -1
}
