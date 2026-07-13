// Command hasprefixfold is generated as a standalone program so the documented HasPrefixFold example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// HasPrefixFold reports whether the string starts with prefix using Unicode simple case folding.
	// An empty prefix is not a match.
	// Similar: HasPrefix and HasSuffixFold.

	// Example: has prefix (case-insensitive)
	v := str.Of("gopher").HasPrefixFold("GO")
	println(v)
	// #bool true
}
