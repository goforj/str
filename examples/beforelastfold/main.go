//go:build ignore
// +build ignore

package main

import "github.com/goforj/str"

func main() {
	// BeforeLastFold returns the substring before the last occurrence of sep using Unicode-aware
	// case-insensitive comparison. If sep is empty or not found, the original string is returned.

	// Example: slice before last separator (case-insensitive)
	v := str.Of("pkg/Path/FILE.txt").BeforeLastFold("/path/").String()
	println(v)
	// #string pkg
}
