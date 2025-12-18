//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Slice returns the substring between rune offsets [start:end).
	// Indices are clamped; if start >= end the result is empty.

	// Example: rune-safe slice
	v := str.Of("naïve café")
	godump.Dump(v.Slice(3, 7).String())
	// #string e ca
}
