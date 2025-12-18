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
	v := str.Of("Go means gophers").ContainsAll("Go", "gopher")
	godump.Dump(v)
	// #bool true
}
