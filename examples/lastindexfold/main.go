//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// LastIndexFold returns the rune index of the last occurrence of sub using Unicode-aware
	// case-insensitive comparison, or -1 if not found.

	// Example: last rune index (case-insensitive)
	v := str.Of("Go gopher GO").LastIndexFold("go")
	println(v)
	// #int 10
}
