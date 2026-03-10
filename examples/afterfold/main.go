//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// AfterFold returns the substring after the first occurrence of sep using Unicode-aware
	// case-insensitive comparison. If sep is empty or not found, the original string is returned.

	// Example: slice after marker (case-insensitive)
	v := str.Of("gopher::GO-team").AfterFold("::go").String()
	println(v)
	// #string -team
}
