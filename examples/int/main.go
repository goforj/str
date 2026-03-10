//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Int parses the string as a base-10 int using strconv.Atoi semantics.

	// Example: parse int
	v, err := str.Of("42").Int()
	println(v, err == nil)
	// #int 42
	// #bool true
}
