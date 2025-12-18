package str

import (
	"strings"
	"unicode"
)

// Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim whitespace
//
//	val := str.Of("  Laravel  ")
//	godump.Dump(val.Trim("").String())
//
//	// #string Laravel
func (s String) Trim(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.Trim(s.s, cutset)}
}

// TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim left
//
//	val := str.Of("  Laravel  ")
//	godump.Dump(val.TrimLeft("").String())
//
//	// #string Laravel
func (s String) TrimLeft(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimLeftFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimLeft(s.s, cutset)}
}

// TrimRight trims trailing characters. If cutset is empty, trims Unicode whitespace.
// @group Cleanup
//
// Example: trim right
//
//	val := str.Of("  Laravel  ")
//	godump.Dump(val.TrimRight("").String())
//
//	// #string   Laravel
func (s String) TrimRight(cutset string) String {
	if cutset == "" {
		return String{s: strings.TrimRightFunc(s.s, unicode.IsSpace)}
	}
	return String{s: strings.TrimRight(s.s, cutset)}
}
