//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// BetweenFirst returns the substring between the first occurrence of start and the first occurrence of end after it.
	// Returns an empty string if markers are missing.

	// Example: minimal span between markers
	val := str.Of("[a] bc [d]")
	godump.Dump(val.BetweenFirst("[", "]").String())

	// #string a
}
