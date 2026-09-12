package str

import "strings"

// ReplacePrefix replaces old with repl when old is a prefix of the string.
// An empty old inserts repl at the beginning.
// Similar: ReplaceSuffix and TrimPrefix.
// @group Replace
//
// Example: replace prefix
//
//	v := str.Of("prefix-value").ReplacePrefix("prefix-", "new-").String()
//	println(v)
//	// #string new-value
func (s String) ReplacePrefix(old, repl string) String {
	if !strings.HasPrefix(s.s, old) {
		return s
	}
	return String{s: repl + s.s[len(old):]}
}
