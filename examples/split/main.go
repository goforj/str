//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Split splits the string by the given separator.

	// Example: split on comma
	v := str.Of("a,b,c").Split(",")
	godump.Dump(v)
	// #[]string [a b c]
}
