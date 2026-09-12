// Command slice is generated as a standalone program so the documented Slice example can be run directly.
package main

import "github.com/goforj/str/v3"

// main keeps this generated example directly runnable with go run.
func main() {
	// Slice returns the substring between rune offsets [start:end).
	// Index and LastIndex return bytes, which must not be passed directly here.
	// Indices are clamped; if start >= end the result is empty.

	// Example: rune-safe slice
	v := str.Of("naïve café").Slice(3, 7).String()
	println(v)
	// #string ve c
}
