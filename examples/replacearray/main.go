//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ReplaceArray replaces all occurrences of each old in olds with repl.

	// Example: replace many
	v := str.Of("The---Laravel---Framework")
	godump.Dump(v.ReplaceArray([]string{"---"}, "-").String())
	// #string The-Laravel-Framework
}
