//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// PadLeft pads the string on the left to reach length runes using pad (defaults to space).

	// Example: pad left
	v := str.Of("go")
	godump.Dump(v.PadLeft(5, " ").String())
	// #string \u00a0\u00a0\u00a0go
}
