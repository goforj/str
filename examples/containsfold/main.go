//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ContainsFold reports whether the string contains any of the provided substrings, using Unicode
	// case folding for comparisons.
	// Empty substrings return true.

	// Example: contains any (case-insensitive)
	val := str.Of("Go means gophers")
	godump.Dump(val.ContainsFold("GOPHER", "rust"))

	// #bool true
}
