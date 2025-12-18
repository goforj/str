package str

import "strings"

// Before returns the substring before the first occurrence of sep.
// If sep is empty or not found, the original string is returned.
// @group Substrings
//
// Example: slice before marker
//
//	val := str.Of("gopher::go")
//	godump.Dump(val.Before("::").String())
//
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
//	val := str.Of("pkg/path/file.txt")
//	godump.Dump(val.BeforeLast("/").String())
//
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
