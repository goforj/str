//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ChopEnd removes the first matching suffix if present.

	// Example: chop end
	val := str.Of("file.txt")
	godump.Dump(val.ChopEnd(".txt", ".md").String())

	// #string file
}
