//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Equals reports whether the string exactly matches other (case-sensitive).

	// Example: exact match
	val := str.Of("gopher")
	godump.Dump(val.Equals("gopher"))

	// #bool true
}
