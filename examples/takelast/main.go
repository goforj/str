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
	val := str.Of("gophers")
	godump.Dump(val.TakeLast(4).String())

	// #string hers
}
