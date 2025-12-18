//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// AfterLast returns the substring after the last occurrence of sep.
	// If sep is empty or not found, the original string is returned.

	// Example: slice after last separator
	val := str.Of("pkg/path/file.txt")
	godump.Dump(val.AfterLast("/").String())

	// #string file.txt
}
