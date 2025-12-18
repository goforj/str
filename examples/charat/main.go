//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// CharAt returns the rune at the given index and true if within bounds.

	// Example: rune lookup
	val := str.Of("gopher")
	r, ok := val.CharAt(2)
	godump.Dump(string(r), ok)

	// #string p
	// #bool true
}
