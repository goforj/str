//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// StartsWith reports whether the string starts with any of the provided prefixes (case-sensitive).

	// Example: starts with any
	v := str.Of("gopher")
	godump.Dump(v.StartsWith("go", "rust"))
	// #bool true
}
