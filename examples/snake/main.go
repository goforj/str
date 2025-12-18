//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Snake converts the string to snake_case using the provided separator (default "_").

	// Example: snake case
	val := str.Of("fooBar baz")
	godump.Dump(val.Snake("_").String())

	// #string foo_bar_baz
}
