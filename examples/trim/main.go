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
	val := str.Of("  Laravel  ")
	godump.Dump(val.Trim("").String())

	// #string Laravel
}
