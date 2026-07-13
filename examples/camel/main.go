//go:build ignore
// +build ignore

// Command camel is generated as a standalone program so the documented Camel example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Camel converts the string to camelCase.
	// Similar: Pascal.

	// Example: camel case
	v := str.Of("foo_bar baz").Camel().String()
	println(v)
	// #string fooBarBaz
}
