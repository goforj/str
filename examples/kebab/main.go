//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Kebab converts the string to kebab-case.

	// Example: kebab case
	v := str.Of("fooBar baz")
	godump.Dump(v.Kebab().String())
	// #string foo-bar-baz
}
