//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// HasSurrounding reports whether the string starts with before and ends with after.
	// If after is empty, before is used for both sides.

	// Example: has matching delimiters
	v := str.Of(`"GoForj"`).HasSurrounding(`"`, "")
	println(v)
	// #bool true
}
