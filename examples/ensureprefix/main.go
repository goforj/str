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
	v := str.Of("path/to").EnsurePrefix("/").String()
	godump.Dump(v)
	// #string /path/to
}
