// Command hassuffix is generated as a standalone program so the documented HasSuffix example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// HasSuffix reports whether the string ends with suffix using a case-sensitive comparison.
	// An empty suffix always matches.
	// Similar: HasSuffixFold and HasPrefix.

	// Example: has suffix
	v := str.Of("gopher").HasSuffix("her")
	println(v)
	// #bool true
}
