//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// IsAlpha reports whether the string contains at least one rune and every rune is a Unicode letter.

	// Example: alphabetic check
	v := str.Of("Gopher").IsAlpha()
	println(v)
	// #bool true
}
