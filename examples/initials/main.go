//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Initials returns the uppercase first rune of each detected word.
	// Words are split the same way as SplitWords, including camelCase boundaries.

	// Example: collect word initials
	v := str.Of("portableNetwork graphics").Initials().String()
	println(v)
	// #string PNG
}
