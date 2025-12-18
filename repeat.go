package str

import "strings"

// Repeat repeats the string count times (non-negative).
// @group Transform
//
// Example: repeat string
//
//	v := str.Of("go")
//	godump.Dump(v.Repeat(3).String())
//	// #string gogogo
func (s String) Repeat(count int) String {
	if count <= 0 {
		return String{s: ""}
	}
	return String{s: strings.Repeat(s.s, count)}
}
