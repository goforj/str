//go:build ignore
// +build ignore

// Command limit is generated as a standalone program so the documented Limit example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// Limit truncates the string to length runes, appending suffix if truncation occurs.

	// Example: limit with suffix
	v := str.Of("Perfectly balanced, as all things should be.").Limit(10, "...").String()
	println(v)
	// #string Perfectly\u0020...
}
