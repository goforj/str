//go:build ignore
// +build ignore

// Command charat is generated as a standalone program so the documented CharAt example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// CharAt returns the rune at the given index and true if within bounds.
	// Similar: Slice and RuneCount.

	// Example: rune lookup
	v, ok := str.Of("gopher").CharAt(2)
	println(string(v), ok)
	// #string p
	// #bool true
}
