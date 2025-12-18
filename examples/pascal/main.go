//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Pascal converts the string to PascalCase.

	// Example: pascal case
	val := str.Of("foo_bar baz")
	godump.Dump(val.Pascal().String())

	// #string FooBarBaz
}
