//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// ChopStart removes the first matching prefix if present.

	// Example: chop start
	val := str.Of("https://goforj.dev")
	godump.Dump(val.ChopStart("https://", "http://").String())

	// #string goforj.dev
}
