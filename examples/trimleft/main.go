//go:build ignore
// +build ignore

// Command trimleft is generated as a standalone program so the documented TrimLeft example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimLeft removes leading Unicode whitespace.
	// Similar: Trim and TrimRight.

	// Example: trim left
	v := str.Of("  GoForj  ").TrimLeft().String()
	println(v)
	// #string GoForj\u0020\u0020
}
