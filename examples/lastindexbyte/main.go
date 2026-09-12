// Command lastindexbyte is generated as a standalone program so the documented LastIndexByte example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// LastIndexByte returns the byte offset of the last byte equal to c, or -1 if absent.

	// Example: LastIndexByte
	v := str.Of("go").LastIndexByte('o')
	println(v)
	// #int 1
}
