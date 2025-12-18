//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Mask replaces the middle of the string with the given rune, revealing revealLeft runes
	// at the start and revealRight runes at the end. Negative reveal values count from the end.
	// If the reveal counts cover the whole string, the original string is returned.

	// Example: mask email
	v := str.Of("gopher@example.com").Mask('*', 3, 4).String()
	godump.Dump(v)
	// #string gop***********.com
}
