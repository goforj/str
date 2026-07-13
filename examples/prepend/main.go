//go:build ignore
// +build ignore

// Command prepend is generated as a standalone program so the documented Prepend example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Prepend concatenates the provided parts to the beginning of the string.
	// Similar: Append.

	// Example: prepend text
	v := str.Of("World").Prepend("Hello ", "Go ").String()
	println(v)
	// #string Hello Go World
}
