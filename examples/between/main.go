//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Between returns the substring between the first occurrence of start and the last occurrence of end.
	// Returns an empty string if either marker is missing or overlapping.

	// Example: between first and last
	v := str.Of("This is my name").Between("This", "name").String()
	godump.Dump(v)
	// #string  is my
}
