//go:build ignore
// +build ignore

// Command append is generated as a standalone program so the documented Append example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Append concatenates the provided parts to the end of the string.
	// Similar: Prepend.

	// Example: append text
	v := str.Of("Go").Append("Forj", "!").String()
	println(v)
	// #string GoForj!
}
