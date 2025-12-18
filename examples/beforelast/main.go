//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// BeforeLast returns the substring before the last occurrence of sep.
	// If sep is empty or not found, the original string is returned.

	// Example: slice before last separator
	val := str.Of("pkg/path/file.txt")
	godump.Dump(val.BeforeLast("/").String())

	// #string pkg/path
}
