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
	v := str.Of("file.txt").ChopEnd(".txt", ".md").String()
	godump.Dump(v)
	// #string file
}
