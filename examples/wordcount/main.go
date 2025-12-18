//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// WordCount returns the number of word tokens (letters/digits runs).

	// Example: count words
	v := str.Of("Hello, world!")
	godump.Dump(v.WordCount())
	// #int 2
}
