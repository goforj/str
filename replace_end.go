package str

import "strings"

// ReplaceSuffix replaces old with repl when old is a suffix of the string.
// Similar: ReplacePrefix and TrimSuffix.
// @group Replace
//
// Example: replace suffix
//
//	v := str.Of("file.old").ReplaceSuffix(".old", ".new").String()
//	println(v)
//	// #string file.new
func (s String) ReplaceSuffix(old, repl string) String {
	if old == "" || !strings.HasSuffix(s.s, old) {
		return s
	}
	return String{s: s.s[:len(s.s)-len(old)] + repl}
}
