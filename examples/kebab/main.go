//go:build ignore
// +build ignore

// Command kebab is generated as a standalone program so the documented Kebab example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Kebab converts the string to kebab-case.
	// Similar: Snake.

	// Example: kebab case
	v := str.Of("fooBar baz").Kebab().String()
	println(v)
	// #string foo-bar-baz
}
