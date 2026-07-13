package str

import "strings"

// Between returns the substring between the first start marker and the first end marker after it.
// It returns an empty string when either marker is empty or missing.
// @group Substrings
//
// Example: between markers
//
//	v := str.Of("[first] and [second]").Between("[", "]").String()
//	println(v)
//	// #string first
func (s String) Between(start, end string) String {
	if start == "" || end == "" {
		return String{s: ""}
	}
	startIdx := strings.Index(s.s, start)
	if startIdx == -1 {
		return String{s: ""}
	}
	startEnd := startIdx + len(start)
	endIdx := strings.Index(s.s[startEnd:], end)
	if endIdx == -1 {
		return String{s: ""}
	}
	return String{s: s.s[startEnd : startEnd+endIdx]}
}
