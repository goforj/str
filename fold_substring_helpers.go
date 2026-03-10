package str

import "unicode"

func foldMatchRange(s, sub string, last bool) (int, int, bool) {
	if sub == "" {
		return 0, 0, false
	}

	haystack := []rune(s)
	needle := []rune(sub)
	if len(needle) > len(haystack) {
		return 0, 0, false
	}

	offsets := make([]int, 0, len(haystack)+1)
	for i := range s {
		offsets = append(offsets, i)
	}
	offsets = append(offsets, len(s))

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
				return offsets[i], offsets[i+len(needle)], true
			}
		}
		return 0, 0, false
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if matchesAt(i) {
			return offsets[i], offsets[i+len(needle)], true
		}
	}

	return 0, 0, false
}
