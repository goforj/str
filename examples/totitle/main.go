//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ToTitle returns a title-cased copy where all letters are mapped using Unicode title case.

	// Example: title map runes
	val := str.Of("ß")
	godump.Dump(val.ToTitle().String())

	// #string SS
}
