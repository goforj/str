//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).

	// Example: ends with any
	val := str.Of("gopher")
	godump.Dump(val.EndsWith("her", "cat"))

	// #bool true
}
