//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EndsWith reports whether the string ends with any of the provided suffixes (case-sensitive).

	// Example: ends with any
	v := str.Of("gopher").EndsWith("her", "cat")
	godump.Dump(v)
	// #bool true
}
