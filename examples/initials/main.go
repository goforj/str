// Command initials is generated as a standalone program so the documented Initials example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Initials returns the uppercase first rune of each detected word.
	// Words are split the same way as SplitWords, including camel case and acronym boundaries.
	// Similar: SplitWords.

	// Example: collect word initials
	v := str.Of("portableNetwork graphics").Initials().String()
	println(v)
	// #string PNG
}
