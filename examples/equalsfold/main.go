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
	v := str.Of("gopher").EqualsFold("GOPHER")
	godump.Dump(v)
	// #bool true
}
