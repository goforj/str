//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// DoesntStartWith reports true if the string starts with none of the provided prefixes (case-sensitive).

	// Example: doesn't start with any
	v := str.Of("gopher")
	godump.Dump(v.DoesntStartWith("go", "rust"))
	// #bool false
}
