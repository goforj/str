//go:build ignore
// +build ignore

// Command remove is generated as a standalone program so the documented Remove example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Remove deletes all occurrences of provided substrings.

	// Example: remove substrings
	v := str.Of("The Go Toolkit").Remove("Go ").String()
	println(v)
	// #string The Toolkit
}
