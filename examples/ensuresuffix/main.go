//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// EnsureSuffix ensures the string ends with suffix, adding it if missing.

	// Example: ensure suffix
	val := str.Of("path/to")
	godump.Dump(val.EnsureSuffix("/").String())

	// #string path/to/
}
