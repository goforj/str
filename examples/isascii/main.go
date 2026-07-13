//go:build ignore
// +build ignore

// Command isascii is generated as a standalone program so the documented IsASCII example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsASCII reports whether the string consists solely of 7-bit ASCII runes.

	// Example: ASCII check
	v := str.Of("gopher").IsASCII()
	println(v)
	// #bool true
}
