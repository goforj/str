//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ContainsAll reports whether the string contains all provided substrings (case-sensitive).
	// Empty substrings are ignored.

	// Example: contains all
	val := str.Of("Go means gophers")
	godump.Dump(val.ContainsAll("Go", "gopher"))

	// #bool true
}
