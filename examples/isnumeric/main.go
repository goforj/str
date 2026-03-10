//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// IsNumeric reports whether the string contains at least one rune and every rune is a Unicode number.

	// Example: numeric check
	v := str.Of("12345").IsNumeric()
	println(v)
	// #bool true
}
