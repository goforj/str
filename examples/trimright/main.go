//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// TrimRight trims trailing characters. If cutset is empty, trims Unicode whitespace.

	// Example: trim right
	val := str.Of("  Laravel  ")
	godump.Dump(val.TrimRight("").String())

	// #string   Laravel
}
