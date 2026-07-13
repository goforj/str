// Command words is generated as a standalone program so the documented Words example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Words limits the string to count words, preserving the source through the
	// selected word boundary and appending suffix if truncated.
	// Similar: SplitWords and WrapWords.

	// Example: limit words
	v := str.Of("Perfectly balanced, as all things should be.").Words(3, " >>>").String()
	println(v)
	// #string Perfectly balanced, as >>>
}
