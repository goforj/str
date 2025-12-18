//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Contains reports whether the string contains any of the provided substrings (case-sensitive).
	// Empty substrings return true to match strings.Contains semantics.

	// Example: contains any
	v := str.Of("Go means gophers").Contains("rust", "gopher")
	godump.Dump(v)
	// #bool true
}
