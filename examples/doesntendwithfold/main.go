//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntEndWithFold reports true if the string ends with none of the provided suffixes using Unicode case folding.

	// Example: doesn't end with (case-insensitive)
	v := str.Of("gopher")
	godump.Dump(v.DoesntEndWithFold("HER"))
	// #bool false
}
