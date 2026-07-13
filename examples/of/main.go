//go:build ignore
// +build ignore

// Command of is generated as a standalone program so the documented Of example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Of wraps a raw string with fluent helpers.

	// Example: wrap raw string
	v := str.Of("gopher")
	println(v.String())
	// #string gopher
}
