//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EnsurePrefix ensures the string starts with prefix, adding it if missing.

	// Example: ensure prefix
	val := str.Of("path/to")
	godump.Dump(val.EnsurePrefix("/").String())

	// #string /path/to
}
