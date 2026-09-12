// Command hasprefixfold is generated as a standalone program so the documented HasPrefixFold example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// HasPrefixFold reports whether the string starts with prefix using Unicode simple case folding.
	// An empty prefix always matches.
	// Similar: HasPrefix and HasSuffixFold.

	// Example: has prefix (case-insensitive)
	v := str.Of("gopher").HasPrefixFold("GO")
	println(v)
	// #bool true
}
