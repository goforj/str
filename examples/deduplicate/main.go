//go:build ignore
// +build ignore

// Command deduplicate is generated as a standalone program so the documented Deduplicate example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Deduplicate collapses consecutive instances of char into a single instance.
	// If char is zero, space is used.
	// Similar: NormalizeSpace.

	// Example: collapse spaces
	v := str.Of("The   Go   Playground").Deduplicate(' ').String()
	println(v)
	// #string The Go Playground
}
