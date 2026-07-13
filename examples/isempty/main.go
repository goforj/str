//go:build ignore
// +build ignore

// Command isempty is generated as a standalone program so the documented IsEmpty example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsEmpty reports whether the string has zero length.
	// Similar: IsBlank.

	// Example: empty check
	v := str.Of("").IsEmpty()
	println(v)
	// #bool true
}
