//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Unwrap removes matching before and after strings if present.

	// Example: unwrap string
	val := str.Of("\"Laravel\"")
	godump.Dump(val.Unwrap("\"", "\"").String())

	// #string Laravel
}
