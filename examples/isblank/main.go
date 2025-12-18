//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// IsBlank reports whether the string contains only Unicode whitespace.

	// Example: blank check
	val := str.Of("  \\t\\n")
	godump.Dump(val.IsBlank())

	// #bool true
}
