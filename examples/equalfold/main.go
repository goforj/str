// Command equalfold is generated as a standalone program so the documented EqualFold example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// EqualFold reports whether the string matches other using Unicode simple case folding.

	// Example: case-insensitive match
	v := str.Of("gopher").EqualFold("GOPHER")
	println(v)
	// #bool true
}
