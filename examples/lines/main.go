//go:build ignore
// +build ignore

// Command lines is generated as a standalone program so the documented Lines example can be run directly.
package main

import (
	"fmt"

	"github.com/goforj/str/v2"
)

// main keeps this generated example directly runnable with go run.
func main() {
	// Lines splits the string into lines after normalizing newline variants.
	// Similar: NormalizeNewlines.

	// Example: split lines
	v := str.Of("a\r\nb\nc").Lines()
	fmt.Println(v)
	// #[]string [a b c]
}
