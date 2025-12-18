//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// WrapWords wraps the string to the given width on word boundaries, using breakStr between lines.

	// Example: wrap words
	v := str.Of("The quick brown fox jumped over the lazy dog.")
	godump.Dump(v.WrapWords(20, "\n").String())
	// #string The quick brown fox\njumped over the lazy\ndog.
}
