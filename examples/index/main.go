// Command index is generated as a standalone program so the documented Index example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Index returns the byte index of the first occurrence of sub, or -1 if not found.
	// Similar: LastIndex.

	// Example: first byte index
	v := str.Of("héllo").Index("llo")
	println(v)
	// #int 3
}
