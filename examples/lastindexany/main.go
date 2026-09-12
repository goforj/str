// Command lastindexany is generated as a standalone program so the documented LastIndexAny example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// LastIndexAny returns the byte offset of the last rune in chars, or -1 if absent.

	// Example: LastIndexAny
	v := str.Of("go").LastIndexAny("o")
	println(v)
	// #int 1
}
