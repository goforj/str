// Command indexrune is generated as a standalone program so the documented IndexRune example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// IndexRune returns the byte offset of the first rune equal to r, or -1 if absent.

	// Example: IndexRune
	v := str.Of("go").IndexRune('o')
	println(v)
	// #int 1
}
