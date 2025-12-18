//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// LastIndex returns the rune index of the last occurrence of sub, or -1 if not found.

	// Example: last rune index
	v := str.Of("go gophers go")
	godump.Dump(v.LastIndex("go"))
	// #int 10
}
