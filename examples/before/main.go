//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Before returns the substring before the first occurrence of sep.
	// If sep is empty or not found, the original string is returned.

	// Example: slice before marker
	val := str.Of("gopher::go")
	godump.Dump(val.Before("::").String())

	// #string gopher
}
