// Command trimspace is generated as a standalone program so the documented TrimSpace example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// TrimSpace removes leading and trailing Unicode whitespace.

	// Example: TrimSpace
	v := str.Of("  GoForj  ").TrimSpace().String()
	println(v)
	// #string GoForj
}
