//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Count returns the number of non-overlapping occurrences of sub.

	// Example: count substring
	v := str.Of("gogophergo").Count("go")
	godump.Dump(v)
	// #int 3
}
