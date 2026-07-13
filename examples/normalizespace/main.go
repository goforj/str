// Command normalizespace is generated as a standalone program so the documented NormalizeSpace example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// NormalizeSpace removes surrounding whitespace and collapses internal whitespace to single spaces.
	// Similar: Trim.

	// Example: normalize whitespace
	v := str.Of("  go   forj  ").NormalizeSpace().String()
	println(v)
	// #string go forj
}
