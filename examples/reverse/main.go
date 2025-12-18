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
	v := str.Of("naïve")
	godump.Dump(v.Reverse().String())
	// #string evïan
}
