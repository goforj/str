//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// IsASCII reports whether the string consists solely of 7-bit ASCII runes.

	// Example: ASCII check
	v := str.Of("gopher").IsASCII()
	godump.Dump(v)
	// #bool true
}
