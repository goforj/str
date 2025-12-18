//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Append concatenates the provided parts to the end of the string.

	// Example: append text
	v := str.Of("Go").Append("Forj", "!").String()
	godump.Dump(v)
	// #string GoForj!
}
