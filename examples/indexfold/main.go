//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// IndexFold returns the rune index of the first occurrence of sub using Unicode-aware
	// case-insensitive comparison, or -1 if not found.

	// Example: first rune index (case-insensitive)
	v := str.Of("Go gopher GO").IndexFold("go")
	println(v)
	// #int 0
}
