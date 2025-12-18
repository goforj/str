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
	v := str.Of("Go means gophers").ContainsAllFold("go", "GOPHER")
	godump.Dump(v)
	// #bool true
}
