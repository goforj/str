//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// CountFold returns the number of non-overlapping occurrences of sub using Unicode-aware
	// case-insensitive comparison.

	// Example: count substring (case-insensitive)
	v := str.Of("GoGOgophergo").CountFold("go")
	println(v)
	// #int 4
}
