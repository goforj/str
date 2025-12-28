package str

import (
	"strings"
	"unicode"
)

// Trim trims leading and trailing characters. If cutset is the zero value (empty string), trims Unicode whitespace.
// @group Cleanup
//
// Example: trim whitespace
//
//	v := str.Of("  GoForj  ").Trim("").String()
//	println(v)
//	// #string GoForj
func (s String) Trim(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.Trim(s.s, cutset)}
}

// TrimSpace trims leading and trailing Unicode whitespace.
// @group Cleanup
//
// Example: trim space
//
//	v := str.Of("  GoForj  ").TrimSpace().String()
//	println(v)
//	// #string GoForj
func (s String) TrimSpace() String {
	if s.s == "" {
		return s
	}
	if trimmed := strings.TrimFunc(s.s, unicode.IsSpace); trimmed == s.s {
		return s
	} else {
		return String{s: trimmed}
	}
}

// TrimLeft trims leading characters. If cutset is the zero value (empty string), trims Unicode whitespace.
// @group Cleanup
//
// Example: trim left
//
//	v := str.Of("  GoForj  ").TrimLeft("").String()
//	println(v)
//	// #string GoForj
func (s String) TrimLeft(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimLeftFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimLeft(s.s, cutset)}
}

// TrimRight trims trailing characters. If cutset is the zero value (empty string), trims Unicode whitespace.
// @group Cleanup
//
// Example: trim right
//
//	v := str.Of("  GoForj  ").TrimRight("").String()
//	println(v)
//	// #string   GoForj
func (s String) TrimRight(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimRightFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimRight(s.s, cutset)}
}
