package str

import "unicode/utf8"

// Len returns the number of runes in the string.
// @group Length
//
// Example: count runes instead of bytes
//
//	val := str.Of("gophers 🦫")
//	godump.Dump(val.Len())
//
//	// #int 9
func (s String) Len() int {
	return utf8.RuneCountInString(s.s)
}

// RuneCount is an alias for Len to make intent explicit in mixed codebases.
// @group Length
//
// Example: alias for Len
//
//	val := str.Of("naïve")
//	godump.Dump(val.RuneCount())
//
//	// #int 5
func (s String) RuneCount() int {
	return s.Len()
}
