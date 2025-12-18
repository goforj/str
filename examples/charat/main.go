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
	v, ok := str.Of("gopher").CharAt(2)
	godump.Dump(string(v), ok)
	// #string p
	// #bool true
}
