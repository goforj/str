//go:build ignore
// +build ignore

// Command excerpt is generated as a standalone program so the documented Excerpt example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Excerpt returns a snippet around the first occurrence of needle with the given radius.
	// If needle is not found, an empty string is returned. If radius <= 0, a default of 100 is used.
	// Omission is used at the start/end when text is trimmed (default "...").

	// Example: excerpt with radius
	v := str.Of("This is my name").Excerpt("my", 3, "...")
	println(v.String())
	// #string ...is my na...
}
