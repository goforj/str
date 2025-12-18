//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// StartsWithFold reports whether the string starts with any of the provided prefixes using Unicode case folding.

	// Example: starts with (case-insensitive)
	v := str.Of("gopher")
	godump.Dump(v.StartsWithFold("GO"))
	// #bool true
}
