package str

import (
	"iter"
	"strings"
)

// Lines returns a single-use iterator over newline-terminated lines.
// Newline bytes are retained; empty input yields no lines and a trailing newline
// does not produce an extra empty line. Use NormalizeNewlines().Split("\n")
// when normalized, delimiter-free fields are wanted.
// @group Split
//
// Example: Lines
//
//	v := slices.Collect(str.Of("a\nb").Lines())
//	fmt.Printf("%q\n", v)
//	// #[]string ["a\\n" "b"]
func (s String) Lines() iter.Seq[string] {
	return strings.Lines(s.s)
}
