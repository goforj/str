// Command clone is generated as a standalone program so the documented Clone example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Clone copies the underlying bytes into a fresh allocation.
	// Use it to release a large backing string retained by a small substring.
	// Empty input returns an empty string without allocating.

	// Example: Clone
	v := str.Of("gopher").Clone().String()
	println(v)
	// #string gopher
}
