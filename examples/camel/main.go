//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/godump"
	"github.com/goforj/str"
)

func main() {
	// Camel converts the string to camelCase.

	// Example: camel case
	val := str.Of("foo_bar baz")
	godump.Dump(val.Camel().String())

	// #string fooBarBaz
}
