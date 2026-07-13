package str

import "strings"

// Lines splits the string into lines after normalizing newline variants.
// Similar: NormalizeNewlines.
// @group Split
//
// Example: split lines
//
//	v := str.Of("a\r\nb\nc").Lines()
//	fmt.Println(v)
//	// #[]string [a b c]
func (s String) Lines() []string {
	normalized := s.NormalizeNewlines().String()
	return strings.Split(normalized, "\n")
}
