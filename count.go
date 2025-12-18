package str

import "strings"

// Count returns the number of non-overlapping occurrences of sub.
// @group Search
//
// Example: count substring
//
//	v := str.Of("gogophergo").Count("go")
//	godump.Dump(v)
//	// #int 3
func (s String) Count(sub string) int {
	if sub == "" {
		return 0
	}
	count := 0
	remain := s.s
	for {
		idx := strings.Index(remain, sub)
		if idx == -1 {
			break
		}
		count++
		remain = remain[idx+len(sub):]
	}
	return count
}
