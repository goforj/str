//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntContainFold reports true if none of the substrings are found using Unicode case folding.
	// Empty substrings are ignored.

	// Example: doesn't contain (case-insensitive)
	v := str.Of("gophers are great")
	godump.Dump(v.DoesntContainFold("GOPHER"))
	// #bool false
}
