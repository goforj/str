//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntStartWithFold reports true if the string starts with none of the provided prefixes using Unicode case folding.

	// Example: doesn't start with (case-insensitive)
	v := str.Of("gopher")
	godump.Dump(v.DoesntStartWithFold("GO"))
	// #bool false
}
