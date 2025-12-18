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
	v := str.Of("https://goforj.dev").ChopStart("https://", "http://").String()
	godump.Dump(v)
	// #string goforj.dev
}
