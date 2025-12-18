//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ReplaceFirst replaces the first occurrence of old with repl.

	// Example: replace first
	val := str.Of("gopher gopher")
	godump.Dump(val.ReplaceFirst("gopher", "go").String())

	// #string go gopher
}
