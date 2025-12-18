//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Take returns the first length runes of the string (clamped).

	// Example: take head
	val := str.Of("gophers")
	godump.Dump(val.Take(3).String())

	// #string gop
}
