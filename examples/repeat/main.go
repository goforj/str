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
	val := str.Of("go")
	godump.Dump(val.Repeat(3).String())

	// #string gogogo
}
