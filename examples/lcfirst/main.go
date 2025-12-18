//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// LcFirst returns the string with the first rune lower-cased.

	// Example: lowercase first rune
	val := str.Of("Gopher")
	godump.Dump(val.LcFirst().String())

	// #string gopher
}
