// Command isalnum is generated as a standalone program so the documented IsAlnum example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsAlnum reports whether the string contains at least one rune and every rune is a Unicode letter or number.

	// Example: alphanumeric check
	v := str.Of("Gopher2025").IsAlnum()
	println(v)
	// #bool true
}
