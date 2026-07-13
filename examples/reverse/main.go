//go:build ignore
// +build ignore

// Command reverse is generated as a standalone program so the documented Reverse example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Reverse returns a rune-safe reversed string.

	// Example: reverse
	v := str.Of("naïve").Reverse().String()
	println(v)
	// #string evïan
}
