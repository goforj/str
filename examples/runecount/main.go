//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// RuneCount is an alias for Len to make intent explicit in mixed codebases.

	// Example: alias for Len
	val := str.Of("naïve")
	godump.Dump(val.RuneCount())

	// #int 5
}
