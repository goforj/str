//go:build ignore
// +build ignore

// Command hasprefix is generated as a standalone program so the documented HasPrefix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// HasPrefix reports whether the string starts with prefix using a case-sensitive comparison.
	// An empty prefix is not a match.
	// Similar: HasPrefixFold and HasSuffix.

	// Example: has prefix
	v := str.Of("gopher").HasPrefix("go")
	println(v)
	// #bool true
}
