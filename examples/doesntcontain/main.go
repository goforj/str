//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntContain reports true if none of the substrings are found (case-sensitive).
	// Empty substrings are ignored.

	// Example: doesn't contain any
	v := str.Of("gophers are great")
	godump.Dump(v.DoesntContain("rust", "beam"))
	// #bool true
}
