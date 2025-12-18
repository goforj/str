//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// IsEmpty reports whether the string has zero length.

	// Example: empty check
	val := str.Of("")
	godump.Dump(val.IsEmpty())

	// #bool true
}
