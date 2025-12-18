//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Limit truncates the string to length runes, appending suffix if truncation occurs.

	// Example: limit with suffix
	val := str.Of("Perfectly balanced, as all things should be.")
	godump.Dump(val.Limit(10, "...").String())

	// #string Perfectly...
}
