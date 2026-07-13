package str

// allRunesMatch reports whether s is non-empty and every rune satisfies match.
func allRunesMatch(s string, match func(rune) bool) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !match(r) {
			return false
		}
	}

	return true
}
