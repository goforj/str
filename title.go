package str

import "strings"

// Title title-cases word-initial letters using strings.Title, preserving other letters.
//
// Deprecated: Like strings.Title, its word boundaries do not handle Unicode
// punctuation properly. Use golang.org/x/text/cases for linguistic title casing.
// @group Case
//
// Example: Title
//
//	v := str.Of("hello WORLD").Title().String()
//	println(v)
//	// #string Hello WORLD
func (s String) Title() String {
	return String{s: strings.Title(s.s)}
}
