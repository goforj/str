//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Index returns the rune index of the first occurrence of sub, or -1 if not found.

	// Example: first rune index
	val := str.Of("héllo")
	godump.Dump(val.Index("llo"))

	// #int 2
}
