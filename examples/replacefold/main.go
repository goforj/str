// Command replacefold is generated as a standalone program so the documented ReplaceFold example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceFold replaces all non-overlapping occurrences of old with repl using Unicode simple case folding.
	// An empty old inserts repl at UTF-8 boundaries, like ReplaceAll.
	// Similar: ReplaceAll.

	// Example: replace all (case-insensitive)
	v := str.Of("go gopher GO").ReplaceFold("GO", "Go").String()
	println(v)
	// #string Go Gopher Go
}
