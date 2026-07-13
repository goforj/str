//go:build ignore
// +build ignore

// Command count is generated as a standalone program so the documented Count example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Count returns the number of non-overlapping occurrences of sub.

	// Example: count substring
	v := str.Of("gogophergo").Count("go")
	println(v)
	// #int 3
}
