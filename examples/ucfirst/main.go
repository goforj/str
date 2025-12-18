//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// UcFirst returns the string with the first rune upper-cased.

	// Example: uppercase first rune
	val := str.Of("gopher")
	godump.Dump(val.UcFirst().String())

	// #string Gopher
}
