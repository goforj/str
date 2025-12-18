//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// TrimLeft trims leading characters. If cutset is empty, trims Unicode whitespace.

	// Example: trim left
	val := str.Of("  Laravel  ")
	godump.Dump(val.TrimLeft("").String())

	// #string Laravel
}
