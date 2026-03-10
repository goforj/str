//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// CommonPrefix returns the longest shared prefix between the string and all provided others.
	// Comparison is rune-safe. If no others are provided, the original string is returned.

	// Example: longest common prefix
	v := str.Of("gopher").CommonPrefix("go", "gold").String()
	println(v)
	// #string go
}
