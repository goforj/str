//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Deduplicate collapses consecutive instances of char into a single instance.
	// If char is zero, space is used.

	// Example: collapse spaces
	val := str.Of("The   Laravel   Framework")
	godump.Dump(val.Deduplicate(' ').String())

	// #string The Laravel Framework
}
