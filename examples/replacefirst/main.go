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
	v := str.Of("gopher gopher")
	godump.Dump(v.ReplaceFirst("gopher", "go").String())
	// #string go gopher
}
