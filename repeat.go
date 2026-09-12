package str

import "strings"

// Repeat repeats the string count times.
// It panics if count is negative or the result length overflows int.
// @group Transform
//
// Example: Repeat
//
//	v := str.Of("go").Repeat(3).String()
//	println(v)
//	// #string gogogo
func (s String) Repeat(count int) String {
	return String{s: strings.Repeat(s.s, count)}
}
