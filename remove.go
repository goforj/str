package str

import "strings"

// Remove deletes all occurrences of provided substrings.
// @group Replace
//
// Example: remove substrings
//
//	val := str.Of("The Laravel Framework")
//	godump.Dump(val.Remove("Laravel ").String())
//
//	// #string The Framework
func (s String) Remove(subs ...string) String {
	out := s.s
	for _, sub := range subs {
		if sub == "" {
			continue
		}
		out = strings.ReplaceAll(out, sub, "")
	}
	return String{s: out}
}
