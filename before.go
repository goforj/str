package str

import "strings"

// Before returns the substring before the first occurrence of sep.
// If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice before marker
//
//	v := str.Of("gopher::go").Before("::").String()
//	godump.Dump(v)
//	// #string gopher
func (s String) Before(sep string) String {
	if sep == "" {
		return s
	}
	idx := strings.Index(s.s, sep)
	if idx == -1 {
		return s
	}
	return String{s: s.s[:idx]}
}

// BeforeLast returns the substring before the last occurrence of sep.
// If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice before last separator
//
//	v := str.Of("pkg/path/file.txt").BeforeLast("/").String()
//	godump.Dump(v)
//	// #string pkg/path
func (s String) BeforeLast(sep string) String {
	if sep == "" {
		return s
	}
	idx := strings.LastIndex(s.s, sep)
	if idx == -1 {
		return s
	}
	return String{s: s.s[:idx]}
}
