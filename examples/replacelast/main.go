//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ReplaceLast replaces the last occurrence of old with repl.

	// Example: replace last
	v := str.Of("gopher gopher")
	godump.Dump(v.ReplaceLast("gopher", "go").String())
	// #string gopher go
}
