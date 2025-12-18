//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// LastWord returns the last word token or empty string.

	// Example: last word
	v := str.Of("Hello world")
	godump.Dump(v.LastWord().String())
	// #string world
}
