// Command runecount is generated as a standalone program so the documented RuneCount example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// RuneCount returns the number of Unicode code points in the string.

	// Example: count runes instead of bytes
	v := str.Of("gophers 🦫").RuneCount()
	println(v)
	// #int 9
}
