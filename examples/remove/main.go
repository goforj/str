//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Remove deletes all occurrences of provided substrings.

	// Example: remove substrings
	v := str.Of("The Go Toolkit")
	godump.Dump(v.Remove("Go ").String())
	// #string The Toolkit
}
