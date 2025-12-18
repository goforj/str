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
	v := str.Of("gopher").EndsWithFold("HER")
	godump.Dump(v)
	// #bool true
}
