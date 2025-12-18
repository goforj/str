//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EqualsFold reports whether the string matches other using Unicode case folding.

	// Example: case-insensitive match
	val := str.Of("gopher")
	godump.Dump(val.EqualsFold("GOPHER"))

	// #bool true
}
