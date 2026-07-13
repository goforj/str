//go:build ignore
// +build ignore

// Command replacefold is generated as a standalone program so the documented ReplaceFold example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceFold replaces all non-overlapping occurrences of old with repl using Unicode simple case folding.
	// An empty old string leaves the receiver unchanged.
	// Similar: ReplaceAll.

	// Example: replace all (case-insensitive)
	v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
	println(v)
	// #string Go Gopher Go
}
