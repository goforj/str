//go:build ignore
// +build ignore

// Command commonprefix is generated as a standalone program so the documented CommonPrefix example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// CommonPrefix returns the longest shared prefix between the string and all provided others.
	// Comparison is rune-safe. If no others are provided, the original string is returned.
	// Similar: CommonSuffix.

	// Example: longest common prefix
	v := str.Of("gopher").CommonPrefix("go", "gold").String()
	println(v)
	// #string go
}
