//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// FirstWord returns the first word token or empty string.

	// Example: first word
	v := str.Of("Hello world")
	godump.Dump(v.FirstWord().String())
	// #string Hello
}
