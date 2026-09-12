// Command replacefirst is generated as a standalone program so the documented ReplaceFirst example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceFirst replaces the first occurrence of old with repl.
	// An empty old inserts repl at the beginning.
	// Similar: ReplaceLast and ReplaceAll.

	// Example: replace first
	v := str.Of("gopher gopher").ReplaceFirst("gopher", "go").String()
	println(v)
	// #string go gopher
}
