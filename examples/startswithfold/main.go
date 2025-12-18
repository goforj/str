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
	val := str.Of("gopher")
	godump.Dump(val.StartsWithFold("GO"))

	// #bool true
}
