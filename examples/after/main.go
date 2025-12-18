//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// After returns the substring after the first occurrence of sep.
	// If sep is empty or not found, the original string is returned.

	// Example: slice after marker
	val := str.Of("gopher::go")
	godump.Dump(val.After("::").String())

	// #string go
}
