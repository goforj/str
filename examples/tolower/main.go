//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ToLower returns a lowercase copy of the string using Unicode rules.

	// Example: lowercase text
	v := str.Of("GoLang")
	godump.Dump(v.ToLower().String())
	// #string golang
}
