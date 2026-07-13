// Command lastword is generated as a standalone program so the documented LastWord example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// LastWord returns the last detected word or an empty string.
	// Similar: FirstWord and SplitWords.

	// Example: last word
	v := str.Of("Hello world").LastWord().String()
	println(v)
	// #string world
}
