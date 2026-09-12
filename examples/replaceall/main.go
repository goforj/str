// Command replaceall is generated as a standalone program so the documented ReplaceAll example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceAll replaces all non-overlapping occurrences of old with new.
	// An empty old matches at the beginning and after each UTF-8 sequence.

	// Example: ReplaceAll
	v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
	println(v)
	// #string Go Gopher Go
}
