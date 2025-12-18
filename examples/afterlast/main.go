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
	v := str.Of("pkg/path/file.txt").AfterLast("/").String()
	godump.Dump(v)
	// #string file.txt
}
