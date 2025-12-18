//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// TakeLast returns the last length runes of the string (clamped).

	// Example: take tail
	v := str.Of("gophers")
	godump.Dump(v.TakeLast(4).String())
	// #string hers
}
