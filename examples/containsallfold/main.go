//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ContainsAllFold reports whether the string contains all provided substrings, using Unicode
	// case folding for comparisons.
	// Empty substrings are ignored.

	// Example: contains all (case-insensitive)
	val := str.Of("Go means gophers")
	godump.Dump(val.ContainsAllFold("go", "GOPHER"))

	// #bool true
}
