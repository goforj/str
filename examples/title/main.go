// Command title is generated as a standalone program so the documented Title example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Title title-cases word-initial letters using strings.Title, preserving other letters.
	//
	// Deprecated: Like strings.Title, its word boundaries do not handle Unicode
	// punctuation properly. Use golang.org/x/text/cases for linguistic title casing.

	// Example: Title
	v := str.Of("hello WORLD").Title().String()
	println(v)
	// #string Hello WORLD
}
