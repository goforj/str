//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// PadRight pads the string on the right to reach length runes using pad (defaults to space).

	// Example: pad right
	val := str.Of("go")
	godump.Dump(val.PadRight(5, ".").String())

	// #string go...
}
