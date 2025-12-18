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
	v := str.Of("path/to").EnsureSuffix("/").String()
	godump.Dump(v)
	// #string path/to/
}
