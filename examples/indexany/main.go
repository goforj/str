// Command indexany is generated as a standalone program so the documented IndexAny example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// IndexAny returns the byte offset of the first rune in chars, or -1 if absent.

	// Example: IndexAny
	v := str.Of("go").IndexAny("o")
	println(v)
	// #int 1
}
