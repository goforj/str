//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// BeforeFold returns the substring before the first occurrence of sep using Unicode-aware
	// case-insensitive comparison. If sep is empty or not found, the original string is returned.

	// Example: slice before marker (case-insensitive)
	v := str.Of("GoPHER::go").BeforeFold("::GO").String()
	println(v)
	// #string GoPHER
}
