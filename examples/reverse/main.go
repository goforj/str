//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Reverse returns a rune-safe reversed string.

	// Example: reverse
	val := str.Of("naïve")
	godump.Dump(val.Reverse().String())

	// #string evïan
}
