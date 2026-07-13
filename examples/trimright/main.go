//go:build ignore
// +build ignore

// Command trimright is generated as a standalone program so the documented TrimRight example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimRight removes trailing Unicode whitespace.
	// Similar: Trim and TrimLeft.

	// Example: trim right
	v := str.Of("  GoForj  ").TrimRight().String()
	println(v)
	// #string \u0020\u0020GoForj
}
