package str

import "strings"

// After returns the substring after the first occurrence of sep.
// If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice after marker
//
//	v := str.Of("gopher::go").After("::").String()
//	godump.Dump(v)
//	// #string go
func (s String) After(sep string) String {
	if sep == "" {
		return s
	}

	idx := strings.Index(s.s, sep)
	if idx == -1 {
		return s
	}

	return String{s: s.s[idx+len(sep):]}
}

// AfterLast returns the substring after the last occurrence of sep.
// If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice after last separator
//
//	v := str.Of("pkg/path/file.txt").AfterLast("/").String()
//	godump.Dump(v)
//	// #string file.txt
func (s String) AfterLast(sep string) String {
	if sep == "" {
		return s
	}

	idx := strings.LastIndex(s.s, sep)
	if idx == -1 {
		return s
	}

	return String{s: s.s[idx+len(sep):]}
}
