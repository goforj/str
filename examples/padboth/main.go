//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// PadBoth pads the string on both sides to reach length runes using pad (defaults to space).

	// Example: pad both
	val := str.Of("go")
	godump.Dump(val.PadBoth(6, "-").String())

	// #string --go--
}
