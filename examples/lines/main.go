// Command lines is generated as a standalone program so the documented Lines example can be run directly.
package main

import (
	"fmt"
	"slices"

	"github.com/goforj/str/v3"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Lines returns a single-use iterator over newline-terminated lines.
	// Newline bytes are retained; empty input yields no lines and a trailing newline
	// does not produce an extra empty line. Use NormalizeNewlines().Split("\n")
	// when normalized, delimiter-free fields are wanted.

	// Example: Lines
	v := slices.Collect(str.Of("a\nb").Lines())
	fmt.Printf("%q\n", v)
	// #[]string ["a\\n" "b"]
}
