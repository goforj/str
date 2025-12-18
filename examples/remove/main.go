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
	val := str.Of("The Laravel Framework")
	godump.Dump(val.Remove("Laravel ").String())

	// #string The Framework
}
