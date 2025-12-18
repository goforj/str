//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Is reports whether the string matches any of the provided wildcard patterns.
	// Patterns use '*' as a wildcard. Case-sensitive.

	// Example: wildcard match
	v := str.Of("foo/bar").Is("foo/*")
	godump.Dump(v)
	// #bool true
}
