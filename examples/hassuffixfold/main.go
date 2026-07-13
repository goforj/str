//go:build ignore
// +build ignore

// Command hassuffixfold is generated as a standalone program so the documented HasSuffixFold example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// HasSuffixFold reports whether the string ends with suffix using Unicode simple case folding.
	// An empty suffix is not a match.
	// Similar: HasSuffix and HasPrefixFold.

	// Example: has suffix (case-insensitive)
	v := str.Of("gopher").HasSuffixFold("HER")
	println(v)
	// #bool true
}
