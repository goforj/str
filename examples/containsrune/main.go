// Command containsrune is generated as a standalone program so the documented ContainsRune example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// ContainsRune reports whether r occurs in the string.

	// Example: ContainsRune
	v := str.Of("café").ContainsRune('é')
	println(v)
	// #bool true
}
