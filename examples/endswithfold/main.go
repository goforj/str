//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EndsWithFold reports whether the string ends with any of the provided suffixes using Unicode case folding.

	// Example: ends with (case-insensitive)
	val := str.Of("gopher")
	godump.Dump(val.EndsWithFold("HER"))

	// #bool true
}
