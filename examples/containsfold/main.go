//go:build ignore
// +build ignore

// Command containsfold is generated as a standalone program so the documented ContainsFold example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// ContainsFold reports whether the string contains sub using Unicode simple case folding.
	// An empty substring is not a match.
	// Similar: Contains.

	// Example: contains substring (case-insensitive)
	v := str.Of("Go means gophers").ContainsFold("GOPHER")
	println(v)
	// #bool true
}
