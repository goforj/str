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
	v := str.Of("")
	godump.Dump(v.IsEmpty())
	// #bool true
}
