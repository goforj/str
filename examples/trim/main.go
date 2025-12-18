//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Trim trims leading and trailing characters. If cutset is empty, trims Unicode whitespace.

	// Example: trim whitespace
	v := str.Of("  GoForj  ")
	godump.Dump(v.Trim("").String())
	// #string GoForj
}
