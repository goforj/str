//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Repeat repeats the string count times (non-negative).

	// Example: repeat string
	v := str.Of("go")
	godump.Dump(v.Repeat(3).String())
	// #string gogogo
}
