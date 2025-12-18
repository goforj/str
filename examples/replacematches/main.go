//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
	"regexp"
)

func main() {
	// ReplaceMatches applies repl to each regex match and returns the result.

	// Example: regex replace with callback
	v := str.Of("Hello 123 World")
	re := regexp.MustCompile(`\d+`)
	godump.Dump(v.ReplaceMatches(re, func(m string) string { return "[" + m + "]" }).String())
	// #string Hello [123] World
}
