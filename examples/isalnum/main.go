//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// IsAlnum reports whether the string contains at least one rune and every rune is a Unicode letter or number.

	// Example: alphanumeric check
	v := str.Of("Gopher2025").IsAlnum()
	println(v)
	// #bool true
}
