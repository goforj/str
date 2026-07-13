// Command lastindex is generated as a standalone program so the documented LastIndex example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.
	// Similar: Index.

	// Example: last rune index
	v := str.Of("go gophers go").LastIndex("go")
	println(v)
	// #int 11
}
