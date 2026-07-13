//go:build ignore
// +build ignore

// Command isalpha is generated as a standalone program so the documented IsAlpha example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// IsAlpha reports whether the string contains at least one rune and every rune is a Unicode letter.

	// Example: alphabetic check
	v := str.Of("Gopher").IsAlpha()
	println(v)
	// #bool true
}
