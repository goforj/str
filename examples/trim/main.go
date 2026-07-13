// Command trim is generated as a standalone program so the documented Trim example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Trim removes leading and trailing Unicode whitespace.
	// Similar: TrimLeft, TrimRight, and TrimChars.

	// Example: trim whitespace
	v := str.Of("  GoForj  ").Trim().String()
	println(v)
	// #string GoForj
}
