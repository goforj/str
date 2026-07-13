//go:build ignore
// +build ignore

// Command takelast is generated as a standalone program so the documented TakeLast example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// TakeLast returns the last length runes of the string (clamped).
	// Similar: Take.

	// Example: take tail
	v := str.Of("gophers").TakeLast(4).String()
	println(v)
	// #string hers
}
