// Command indexbyte is generated as a standalone program so the documented IndexByte example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// IndexByte returns the byte offset of the first byte equal to c, or -1 if absent.

	// Example: IndexByte
	v := str.Of("go").IndexByte('o')
	println(v)
	// #int 1
}
