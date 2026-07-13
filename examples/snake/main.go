//go:build ignore
// +build ignore

// Command snake is generated as a standalone program so the documented Snake example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Snake converts the string to snake_case.
	// Similar: Kebab.

	// Example: snake case
	v := str.Of("fooBar baz").Snake().String()
	println(v)
	// #string foo_bar_baz
}
