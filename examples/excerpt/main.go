//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Excerpt returns a snippet around the first occurrence of needle with the given radius.
	// If needle is not found, an empty string is returned. If radius <= 0, a default of 100 is used.
	// Omission is used at the start/end when text is trimmed (default "...").

	// Example: excerpt with radius
	v := str.Of("This is my name").Excerpt("my", 3, "...")
	godump.Dump(v.String())
	// #string ...is my na...
}
