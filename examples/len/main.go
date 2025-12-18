//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Len returns the number of runes in the string.

	// Example: count runes instead of bytes
	v := str.Of("gophers 🦫").Len()
	godump.Dump(v)
	// #int 9
}
