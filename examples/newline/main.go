//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// NewLine appends a newline character to the string.

	// Example: append newline
	v := str.Of("hello").NewLine().Append("world").String()
	godump.Dump(v)
	// #string hello\nworld
}
