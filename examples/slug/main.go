//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Slug produces an ASCII slug using the provided separator (default "-"), stripping accents where possible.
	// Not locale-aware; intended for identifiers/URLs.

	// Example: build slug
	val := str.Of("Laravel Framework")
	godump.Dump(val.Slug("-").String())

	// #string laravel-framework
}
