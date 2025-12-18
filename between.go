package str

import "strings"

// Between returns the substring between the first occurrence of start and the last occurrence of end.
// Returns an empty string if either marker is missing or overlapping.
// @group Substrings
//
// Example: between first and last
//
//	v := str.Of("This is my name").Between("This", "name").String()
//	godump.Dump(v)
//	// #string  is my
func (s String) Between(start, end string) String {
	if start == "" || end == "" {
		return String{s: ""}
	}
	startIdx := strings.Index(s.s, start)
	endIdx := strings.LastIndex(s.s, end)
	if startIdx == -1 || endIdx == -1 {
		return String{s: ""}
	}
	startEnd := startIdx + len(start)
	if startEnd > endIdx {
		return String{s: ""}
	}
	return String{s: s.s[startEnd:endIdx]}
}

// BetweenFirst returns the substring between the first occurrence of start and the first occurrence of end after it.
// Returns an empty string if markers are missing.
// @group Substrings
//
// Example: minimal span between markers
//
//	v := str.Of("[a] bc [d]").BetweenFirst("[", "]").String()
//	godump.Dump(v)
//	// #string a
func (s String) BetweenFirst(start, end string) String {
	if start == "" || end == "" {
		return String{s: ""}
	}
	startIdx := strings.Index(s.s, start)
	if startIdx == -1 {
		return String{s: ""}
	}
	startEnd := startIdx + len(start)
	rest := s.s[startEnd:]
	endIdx := strings.Index(rest, end)
	if endIdx == -1 {
		return String{s: ""}
	}
	return String{s: rest[:endIdx]}
}
