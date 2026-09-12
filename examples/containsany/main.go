// Command containsany is generated as a standalone program so the documented ContainsAny example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ContainsAny reports whether any rune in chars occurs in the string.

	// Example: ContainsAny
	v := str.Of("gopher").ContainsAny("aeiou")
	println(v)
	// #bool true
}
