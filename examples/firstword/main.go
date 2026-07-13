// Command firstword is generated as a standalone program so the documented FirstWord example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// FirstWord returns the first detected word or an empty string.
	// Similar: LastWord and SplitWords.

	// Example: first word
	v := str.Of("Hello world")
	println(v.FirstWord().String())
	// #string Hello
}
