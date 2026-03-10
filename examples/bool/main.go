//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Bool parses the string as a bool using strconv.ParseBool semantics.

	// Example: parse bool
	v, err := str.Of("true").Bool()
	println(v, err == nil)
	// #bool true
	// #bool true
}
