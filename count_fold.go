package str

// CountFold returns the number of non-overlapping occurrences of sub using Unicode-aware
// case-insensitive comparison.
// @group Search
//
// Example: count substring (case-insensitive)
//
//	v := str.Of("GoGOgophergo").CountFold("go")
//	println(v)
//	// #int 4
func (s String) CountFold(sub string) int {
	if sub == "" {
		return 0
	}

	count := 0
	remain := s.s
	for {
		_, end, ok := foldMatchRange(remain, sub, false)
		if !ok {
			break
		}
		count++
		remain = remain[end:]
	}

	return count
}
