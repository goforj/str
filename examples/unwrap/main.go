//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// Unwrap removes matching before and after strings if present.

	// Example: unwrap string
	v := str.Of(`"GoForj"`).Unwrap(`"`, `"`).String()
	str.Dump(v)
	// #string GoForj
}
