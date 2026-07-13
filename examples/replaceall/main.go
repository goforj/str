//go:build ignore
// +build ignore

// Command replaceall is generated as a standalone program so the documented ReplaceAll example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ReplaceAll replaces all occurrences of old with new in the string.
	// If old is empty, the original string is returned unchanged.

	// Example: replace all occurrences
	v := str.Of("go gopher go").ReplaceAll("go", "Go").String()
	println(v)
	// #string Go Gopher Go
}
