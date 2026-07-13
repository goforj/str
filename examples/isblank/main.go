//go:build ignore
// +build ignore

// Command isblank is generated as a standalone program so the documented IsBlank example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsBlank reports whether the string contains only Unicode whitespace.
	// Similar: IsEmpty.

	// Example: blank check
	v := str.Of("  \t\n")
	println(v.IsBlank())
	// #bool true
}
