//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Words limits the string to count words, appending suffix if truncated.

	// Example: limit words
	val := str.Of("Perfectly balanced, as all things should be.")
	godump.Dump(val.Words(3, " >>>").String())

	// #string Perfectly balanced as >>>
}
