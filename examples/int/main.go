//go:build ignore
// +build ignore

// Command int is generated as a standalone program so the documented Int example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Int parses the string as a base-10 int using strconv.Atoi semantics.
	// Similar: Bool and Float64.

	// Example: parse int
	v, err := str.Of("42").Int()
	println(v, err == nil)
	// #int 42
	// #bool true
}
