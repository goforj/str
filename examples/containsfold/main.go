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
	v := str.Of("Go means gophers").ContainsFold("GOPHER", "rust")
	godump.Dump(v)
	// #bool true
}
