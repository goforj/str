//go:build ignore
// +build ignore

// Command ucfirst is generated as a standalone program so the documented UcFirst example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// UcFirst returns the string with the first rune upper-cased.
	// Similar: LcFirst and ToUpper.

	// Example: uppercase first rune
	v := str.Of("gopher").UcFirst().String()
	println(v)
	// #string Gopher
}
